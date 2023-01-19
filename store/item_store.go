package store

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/megre/dto"
)

type ItemServiceStore interface {
	AddItem(ctx context.Context, item dto.ItemRequest) (itemResponse dto.ItemResponse, err error)
	GetItem(ctx context.Context, itemID string) (itemResponse dto.ItemResponse, err error)
}

type itemServiceStorer struct {
	db *sqlx.DB
}

func NewItemServiceStore(db *sqlx.DB) ItemServiceStore {
	return &itemServiceStorer{
		db: db,
	}
}

func (service *itemServiceStorer) AddItem(ctx context.Context, item dto.ItemRequest) (itemResponse dto.ItemResponse, err error) {
	err = service.db.GetContext(ctx, &itemResponse, addItemQuery, item.ID, item.AccountID, item.Name, item.ImageURL, item.Quantity)
	return
}

func (service *itemServiceStorer) GetItem(ctx context.Context, itemID string) (itemResponse dto.ItemResponse, err error) {
	err = service.db.GetContext(ctx, &itemResponse, getItemByIDQuery, itemID)
	return
}
