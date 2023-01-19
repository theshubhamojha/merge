package items

import (
	"context"
	"database/sql"

	"github.com/megre/dto"
	"github.com/megre/merrors"
	"github.com/megre/store"
	"github.com/megre/utils"
)

type ItemService interface {
	AddItem(ctx context.Context, itemRequest dto.ItemRequest) (itemResponse dto.ItemResponse, err *merrors.Error)
	GetItem(ctx context.Context, itemID string) (itemResponse dto.ItemResponse, err *merrors.Error)
}

type itemService struct {
	itemStoreService store.ItemServiceStore
}

func NewItemService(itemStore store.ItemServiceStore) ItemService {
	return &itemService{
		itemStoreService: itemStore,
	}
}

func (service *itemService) AddItem(ctx context.Context, itemRequest dto.ItemRequest) (itemResponse dto.ItemResponse, err *merrors.Error) {
	itemRequest.ID = utils.GenerateRandomID(ctx, "item_")
	itemRequest.AccountID = ctx.Value(dto.AccountID).(string)

	resp, dbErr := service.itemStoreService.AddItem(ctx, itemRequest)
	if dbErr != nil {
		err = &merrors.Error{
			Message:   "something went wrong",
			ErrorCode: merrors.InternalServerError,
			Error:     dbErr,
		}
		return
	}

	return resp, nil
}

func (service *itemService) GetItem(ctx context.Context, itemID string) (itemResponse dto.ItemResponse, err *merrors.Error) {
	resp, dbErr := service.itemStoreService.GetItem(ctx, itemID)
	if dbErr != nil {
		if dbErr == sql.ErrNoRows {
			err = &merrors.Error{
				Message:   "item doesn't exist in database",
				ErrorCode: merrors.DataNotExist,
			}
			return
		}

		err = &merrors.Error{
			Message:   "unable to get item data",
			ErrorCode: merrors.InternalServerError,
			Error:     dbErr,
		}
		return
	}

	return resp, nil
}
