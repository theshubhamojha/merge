package accounts_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/megre/accounts"
	"github.com/megre/dto"
	"github.com/megre/merrors"
	storeMocks "github.com/megre/store/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite

	jwtSecret        string
	jwtExpiry        int
	accountStoreMock *storeMocks.AccountStore
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (suite *testSuite) SetupTest() {
	suite.jwtExpiry = 100
	suite.jwtSecret = "secret"
	suite.accountStoreMock = &storeMocks.AccountStore{}
}
func (suite *testSuite) TestCreateAccount() {
	t := suite.T()

	t.Run("when account creation is successful", func(t *testing.T) {
		suite.SetupTest()

		createAccountRequest := dto.CreateAccountRequest{
			Account: dto.Account{
				Name:     "shubham",
				Email:    "mailme.shubham.ojha@gmail.com",
				Password: "shubham",
				Role:     "admin",
			},
		}

		suite.accountStoreMock.On("CreateAccount", mock.Anything, mock.Anything).Return(dto.Account{Name: "shubham", Email: "mailme.shubham.ojha@gmail.com", ID: "acc_123", Role: "admin"}, nil)
		suite.accountStoreMock.On("GetAccountDetailsByEmail", mock.Anything, "mailme.shubham.ojha@gmail.com").Return(dto.Account{}, sql.ErrNoRows)
		accountService := accounts.NewAccountService(suite.jwtSecret, suite.jwtExpiry, suite.accountStoreMock)
		response, err := accountService.CreateAccount(context.Background(), createAccountRequest)

		assert.Equal(t, err, nil)
		assert.Equal(t, response.Email, "mailme.shubham.ojha@gmail.com")
	})

	t.Run("when the email already exists", func(t *testing.T) {
		suite.SetupTest()

		createAccountRequest := dto.CreateAccountRequest{
			Account: dto.Account{
				Name:     "shubham",
				Email:    "mailme.shubham.ojha@gmail.com",
				Password: "shubham",
				Role:     "admin",
			},
		}

		suite.accountStoreMock.On("GetAccountDetailsByEmail", mock.Anything, "mailme.shubham.ojha@gmail.com").Return(dto.Account{Email: "mailme.shubham.ojha@gmail.com", Name: "shubham"}, nil)

		accountService := accounts.NewAccountService(suite.jwtSecret, suite.jwtExpiry, suite.accountStoreMock)
		_, err := accountService.CreateAccount(context.Background(), createAccountRequest)

		assert.NotEqual(t, err, nil)
		assert.Equal(t, err.ErrorCode, merrors.EmailAlreadyExist)
	})

	t.Run("when there is a db connection issue", func(t *testing.T) {
		suite.SetupTest()

		createAccountRequest := dto.CreateAccountRequest{
			Account: dto.Account{
				Name:     "shubham",
				Email:    "mailme.shubham.ojha@gmail.com",
				Password: "shubham",
				Role:     "admin",
			},
		}

		suite.accountStoreMock.On("GetAccountDetailsByEmail", mock.Anything, "mailme.shubham.ojha@gmail.com").Return(dto.Account{}, errors.New("connection issue"))

		accountService := accounts.NewAccountService(suite.jwtSecret, suite.jwtExpiry, suite.accountStoreMock)
		_, err := accountService.CreateAccount(context.Background(), createAccountRequest)

		assert.NotEqual(t, err, nil)
		assert.Equal(t, err.ErrorCode, merrors.InternalServerError)
	})
}

func (suite *testSuite) TestSuspendAccount() {
	t := suite.T()

	t.Run("when account suspension is successfull", func(t *testing.T) {
		suite.SetupTest()

		suite.accountStoreMock.On("SuspendAccount", mock.Anything, "acc_123").Return(nil)

		accountService := accounts.NewAccountService(suite.jwtSecret, suite.jwtExpiry, suite.accountStoreMock)
		err := accountService.SuspendAccount(context.Background(), "acc_123")

		assert.Equal(t, err, nil)
	})

	t.Run("when the account doesn't exist in database", func(t *testing.T) {
		suite.SetupTest()

		suite.accountStoreMock.On("SuspendAccount", mock.Anything, "acc_123").Return(sql.ErrNoRows)

		accountService := accounts.NewAccountService(suite.jwtSecret, suite.jwtExpiry, suite.accountStoreMock)
		err := accountService.SuspendAccount(context.Background(), "acc_123")
		assert.NotEqual(t, err, nil)
		assert.Equal(t, err.ErrorCode, merrors.EmailNotExist)
	})

	t.Run("when database is throwing errors", func(t *testing.T) {
		suite.SetupTest()

		suite.accountStoreMock.On("SuspendAccount", mock.Anything, "acc_123").Return(errors.New("some db error"))

		accountService := accounts.NewAccountService(suite.jwtSecret, suite.jwtExpiry, suite.accountStoreMock)
		err := accountService.SuspendAccount(context.Background(), "acc_123")
		assert.NotEqual(t, err, nil)
		assert.Equal(t, err.ErrorCode, merrors.InternalServerError)
	})
}

func (suite *testSuite) TestLoginAccount() {
	t := suite.T()

	t.Run("when account login is successful", func(t *testing.T) {
		suite.SetupTest()

		request := dto.LoginRequest{
			EmailID:  "shubham@gmail.com",
			Password: "password",
		}

		sampleValidHash := "$2a$04$jWpmNTBZlHYFXah3DuBXNeU8i3HU1kZZNDSHn5rc2dhWNrNbuJBqG"
		dbResponse := dto.Account{
			ID:        "acc_123",
			Name:      "shubham",
			Email:     "shubham@gmail.com",
			Password:  sampleValidHash,
			Role:      "admin",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		suite.accountStoreMock.On("GetAccountDetailsByEmail", mock.Anything, "shubham@gmail.com").Return(dbResponse, nil)

		accountService := accounts.NewAccountService(suite.jwtSecret, suite.jwtExpiry, suite.accountStoreMock)
		response, err := accountService.LoginAccount(context.Background(), request)

		assert.Equal(t, err, nil)
		t.Log(response.Token)
		assert.NotEqual(t, len(response.Token), 0)
		assert.Equal(t, response.Email, "shubham@gmail.com")
	})

	t.Run("when account is suspended", func(t *testing.T) {
		suite.SetupTest()

		request := dto.LoginRequest{
			EmailID:  "shubham@gmail.com",
			Password: "password",
		}

		sampleValidHash := "$2a$04$jWpmNTBZlHYFXah3DuBXNeU8i3HU1kZZNDSHn5rc2dhWNrNbuJBqG"
		dbResponse := dto.Account{
			ID:        "acc_123",
			Name:      "shubham",
			Email:     "shubham@gmail.com",
			Password:  sampleValidHash,
			Role:      "admin",
			IsActive:  false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		suite.accountStoreMock.On("GetAccountDetailsByEmail", mock.Anything, "shubham@gmail.com").Return(dbResponse, nil)

		accountService := accounts.NewAccountService(suite.jwtSecret, suite.jwtExpiry, suite.accountStoreMock)
		_, err := accountService.LoginAccount(context.Background(), request)

		assert.NotEqual(t, err, nil)
		assert.Equal(t, err.ErrorCode, merrors.AccountSuspended)
	})

	t.Run("when the password is wrong", func(t *testing.T) {
		suite.SetupTest()

		request := dto.LoginRequest{
			EmailID:  "shubham@gmail.com",
			Password: "wrong_password",
		}

		sampleValidHash := "$2a$04$jWpmNTBZlHYFXah3DuBXNeU8i3HU1kZZNDSHn5rc2dhWNrNbuJBqG"
		dbResponse := dto.Account{
			ID:        "acc_123",
			Name:      "shubham",
			Email:     "shubham@gmail.com",
			Password:  sampleValidHash,
			Role:      "admin",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		suite.accountStoreMock.On("GetAccountDetailsByEmail", mock.Anything, "shubham@gmail.com").Return(dbResponse, nil)

		accountService := accounts.NewAccountService(suite.jwtSecret, suite.jwtExpiry, suite.accountStoreMock)
		response, err := accountService.LoginAccount(context.Background(), request)

		assert.Equal(t, err.ErrorCode, merrors.IncorrectCredentials)
		assert.Equal(t, len(response.Token), 0)
		assert.Equal(t, response.Email, "")
	})

	t.Run("when email provided doesn't exist in database", func(t *testing.T) {
		suite.SetupTest()

		suite.accountStoreMock.On("GetAccountDetailsByEmail", mock.Anything, "shubham@gmail.com").Return(dto.Account{}, sql.ErrNoRows)

		request := dto.LoginRequest{
			EmailID:  "shubham@gmail.com",
			Password: "wrong_password",
		}
		accountService := accounts.NewAccountService(suite.jwtSecret, suite.jwtExpiry, suite.accountStoreMock)
		_, err := accountService.LoginAccount(context.Background(), request)

		assert.NotEqual(t, err, nil)
		assert.Equal(t, err.ErrorCode, merrors.EmailNotExist)
	})

	t.Run("when there is some database issue to get account details", func(t *testing.T) {
		suite.SetupTest()

		suite.accountStoreMock.On("GetAccountDetailsByEmail", mock.Anything, "shubham@gmail.com").Return(dto.Account{}, errors.New("some db issue"))

		request := dto.LoginRequest{
			EmailID:  "shubham@gmail.com",
			Password: "wrong_password",
		}
		accountService := accounts.NewAccountService(suite.jwtSecret, suite.jwtExpiry, suite.accountStoreMock)
		_, err := accountService.LoginAccount(context.Background(), request)

		assert.NotEqual(t, err, nil)
		assert.Equal(t, err.ErrorCode, merrors.InternalServerError)
	})

}
