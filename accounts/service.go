package accounts

import (
	"context"
	"database/sql"

	"github.com/megre/dto"
	"github.com/megre/merrors"
	"github.com/megre/store"
	"github.com/megre/utils"
)

type AccountService interface {
	CreateAccount(ctx context.Context, request dto.CreateAccountRequest) (response dto.CreateAccountResponse, err *merrors.Error)
	SuspendAccount(ctx context.Context, accountId string) (err *merrors.Error)
	LoginAccount(ctx context.Context, loginRequest dto.LoginRequest) (response dto.LoginResponse, err *merrors.Error)
}

type accountService struct {
	jwtSecret string
	jwtExpiry int

	accountStorerService store.AccountStore
}

func NewAccountService(jwtSecret string, jwtExpiry int, accountStorerService store.AccountStore) AccountService {
	return &accountService{
		jwtSecret:            jwtSecret,
		jwtExpiry:            jwtExpiry,
		accountStorerService: accountStorerService,
	}
}

func (service *accountService) CreateAccount(ctx context.Context, request dto.CreateAccountRequest) (response dto.CreateAccountResponse, err *merrors.Error) {
	verifyResponse, verifyError := service.accountStorerService.GetAccountDetailsByEmail(ctx, request.Email)
	if verifyError != nil && verifyError != sql.ErrNoRows {
		err = &merrors.Error{
			Message:   "error creating account",
			ErrorCode: merrors.InternalServerError,
			Error:     verifyError,
		}
		return
	}
	if verifyResponse.Email != "" {
		err = &merrors.Error{
			Message:   "email already exisits",
			ErrorCode: merrors.EmailAlreadyExist,
			Error:     verifyError,
		}
		return
	}

	request.ID = utils.GenerateRandomID(ctx, "acc_")
	request.IsActive = true
	request.Password = utils.HashAndSalt(ctx, []byte(request.Password))

	dbResponse, dbError := service.accountStorerService.CreateAccount(ctx, request)
	if dbError != nil {
		err = &merrors.Error{
			Message:   "error creating account",
			ErrorCode: merrors.InternalServerError,
			Error:     dbError,
		}
		return
	}

	response = dto.CreateAccountResponse{
		Id:        dbResponse.ID,
		Name:      dbResponse.Name,
		Email:     dbResponse.Email,
		Role:      dbResponse.Role,
		CreatedAt: dbResponse.CreatedAt,
		UpdatedAt: dbResponse.UpdatedAt,
	}

	return
}

func (service *accountService) SuspendAccount(ctx context.Context, accountId string) (err *merrors.Error) {
	dbError := service.accountStorerService.SuspendAccount(ctx, accountId)
	if dbError != nil && dbError == sql.ErrNoRows {
		err = &merrors.Error{
			Message:   "the account id doesn't exisit",
			ErrorCode: merrors.EmailNotExist,
		}
		return
	}
	if dbError != nil {
		err = &merrors.Error{
			Message:   "error suspending account",
			ErrorCode: merrors.InternalServerError,
		}
		return
	}

	return
}

func (service *accountService) LoginAccount(ctx context.Context, request dto.LoginRequest) (response dto.LoginResponse, err *merrors.Error) {
	accountDetails, dbError := service.accountStorerService.GetAccountDetailsByEmail(ctx, request.EmailID)
	if dbError == sql.ErrNoRows {
		err = &merrors.Error{
			Message:   "user doesn't exist",
			ErrorCode: merrors.EmailNotExist,
		}
		return
	}
	if dbError != nil {
		err = &merrors.Error{
			Message:   "internal server error",
			ErrorCode: merrors.InternalServerError,
		}
		return
	}

	if !accountDetails.IsActive {
		err = &merrors.Error{
			Message:   "account is suspended",
			ErrorCode: merrors.AccountSuspended,
		}
		return
	}

	isCredentialValid := utils.VerifyHashSalt(ctx, request.Password, accountDetails.Password)
	if !isCredentialValid {
		err = &merrors.Error{
			Message:   "incorrect credentials provided",
			ErrorCode: merrors.IncorrectCredentials,
		}
		return
	}

	token, jwtErr := utils.GenerateJWTToken(accountDetails.Email, accountDetails.Role, accountDetails.ID, service.jwtSecret, service.jwtExpiry)
	if jwtErr != nil {
		err = &merrors.Error{
			Message:   "error logging in user",
			ErrorCode: merrors.InternalServerError,
		}
		return
	}

	response = dto.LoginResponse{
		Token: token,
		Email: accountDetails.Email,
		Name:  accountDetails.Name,
	}

	return
}
