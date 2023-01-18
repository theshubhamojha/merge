package store

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/megre/dto"
)

type AccountStore interface {
	CreateAccount(ctx context.Context, request dto.CreateAccountRequest) (response dto.Account, err error)
	SuspendAccount(ctx context.Context, accountID string) (err error)
	GetAccountDetailsByEmail(ctx context.Context, email string) (response dto.Account, err error)
}

type accountStorer struct {
	db *sqlx.DB
}

func NewAccountStorerService(db *sqlx.DB) AccountStore {
	return &accountStorer{
		db: db,
	}
}

func (service *accountStorer) CreateAccount(ctx context.Context, request dto.CreateAccountRequest) (response dto.Account, err error) {
	err = service.db.GetContext(ctx, &response, createAccountQuery, request.ID, request.Name, request.Email, request.Password, request.Role)
	if err != nil {
		return
	}

	return
}

func (service *accountStorer) SuspendAccount(ctx context.Context, accountID string) (err error) {
	_, err = service.db.ExecContext(ctx, suspendAccountQuery, accountID)
	if err != nil {
		return
	}

	return
}

func (service *accountStorer) GetAccountDetailsByEmail(ctx context.Context, email string) (response dto.Account, err error) {
	err = service.db.GetContext(ctx, &response, getAccountDetailsByEmailQuery, email)
	if err != nil {
		return
	}

	return
}
