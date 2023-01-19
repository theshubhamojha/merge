package dto

import (
	"time"

	"github.com/megre/merrors"
)

type ItemRequest struct {
	ID        string `json:"id"`
	AccountID string `json:"account_id"`
	Name      string `json:"name"`
	ImageURL  string `json:"image_url"`
	Quantity  int    `json:"quantity"`
}

type ItemResponse struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	AccountID string    `json:"account_id" db:"account_id"`
	ImageURL  string    `json:"image_url" db:"image_url"`
	Quantity  int       `json:"quantity" db:"quantity"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (r *ItemRequest) ValidateRequest() (err *merrors.Error) {
	if len(r.Name) == 0 {
		err = &merrors.Error{
			Message:   "item name cannot be empty",
			ErrorCode: merrors.BadRequets,
		}
		return
	}

	if r.Quantity <= 0 {
		err = &merrors.Error{
			Message:   "item quantity cannot be less than or equal to 0",
			ErrorCode: merrors.BadRequets,
		}
		return
	}

	return
}
