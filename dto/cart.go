package dto

import (
	"time"

	"github.com/megre/merrors"
)

type UpsertCartRequest struct {
	ItemId   string         `json:"item_id"`
	Quantity int            `json:"quantity"`
	Type     CartUpdateType `json:"type"`
}

type Cart struct {
	ID        string    `json:"id" db:"id"`
	AccountID string    `json:"account_id" db:"account_id"`
	ItemID    string    `json:"item_id" db:"item_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	IsDeleted bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CartItems struct {
	ID        string `json:"id" db:"id"`
	AccountID string `json:"account_id" db:"account_id"`
	Name      string `json:"name" db:"name"`
	Quantity  int    `json:"quantity" db:"quantity"`
	ImageURL  string `json:"image_url" db:"image_url"`
}

type CartUpdateType string

const (
	AddItem    CartUpdateType = "add"
	RemoveItem CartUpdateType = "remove"
)

func (r *UpsertCartRequest) ValidateRequest() (err *merrors.Error) {
	if len(r.ItemId) == 0 {
		err = &merrors.Error{
			Message:   "missing item id",
			ErrorCode: merrors.BadRequets,
		}
	}

	if r.Quantity == 0 {
		err = &merrors.Error{
			Message:   "quantity cannot be equal to 0",
			ErrorCode: merrors.BadRequets,
		}
		return
	}

	if r.Type != AddItem && r.Type != RemoveItem {
		err = &merrors.Error{
			Message:   "type can only be add/remove",
			ErrorCode: merrors.BadRequets,
		}
		return
	}

	if (r.Type == AddItem && r.Quantity < 0) || (r.Type == RemoveItem && r.Quantity > 0) {
		err = &merrors.Error{
			Message:   "mismatch between type and quantity",
			ErrorCode: merrors.BadRequets,
		}
		return
	}

	return
}
