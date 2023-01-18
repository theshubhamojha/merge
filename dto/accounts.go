package dto

import (
	"net/mail"
	"time"

	"github.com/megre/merrors"
)

type Account struct {
	ID        string    `json:"id,omitempty" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Role      RoleType  `json:"role" db:"role"`
	IsActive  bool      `json:"is_active,omitempty" db:"is_active"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type CreateAccountRequest struct {
	Account
}

type CreateAccountResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      RoleType  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *CreateAccountRequest) ValidateCreateAccountRequest() (err *merrors.Error) {
	if r.Role != Admin && r.Role != User {
		err = &merrors.Error{
			Message:   "invalid role type",
			ErrorCode: merrors.BadRequets,
		}
		return
	}

	if _, mailError := mail.ParseAddress(r.Email); mailError != nil {
		err = &merrors.Error{
			Message:   "invalid email format",
			ErrorCode: merrors.BadRequets,
		}
		return
	}

	if len(r.Password) < 8 {
		err = &merrors.Error{
			Message:   "password length too small, use atleast 8 characters",
			ErrorCode: merrors.BadRequets,
		}
		return
	}

	if len(r.Name) == 0 {
		err = &merrors.Error{
			Message:   "missing value in `name` feild",
			ErrorCode: merrors.BadRequets,
		}
		return
	}
	return
}
