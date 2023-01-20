package store

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/megre/dto"
)

type CartServiceStore interface {
	GetCartDetail(ctx context.Context, itemID string, accountID string) (cart dto.Cart, err error)
	InsertCart(ctx context.Context, cartItem dto.Cart) (response dto.Cart, err error)
	UpdateQuantity(ctx context.Context, cartID string, quantity int) (response dto.Cart, err error)
	ListCartItems(ctx context.Context, accountID string, pageNumber int) (cart []dto.CartItems, err error)
}

type cartServiceStorer struct {
	db          *sqlx.DB
	itemPerPage int
}

func NewCartServiceStore(db *sqlx.DB, itemPerPage int) CartServiceStore {
	return &cartServiceStorer{
		itemPerPage: itemPerPage,
		db:          db,
	}
}

func (service *cartServiceStorer) GetCartDetail(ctx context.Context, itemID string, accountID string) (cart dto.Cart, err error) {
	err = service.db.GetContext(ctx, &cart, getCartDetailsQuery, accountID, itemID)
	return
}

func (service *cartServiceStorer) InsertCart(ctx context.Context, cartItem dto.Cart) (response dto.Cart, err error) {
	err = service.db.GetContext(ctx, &response, insertCartQuery, cartItem.ID, cartItem.AccountID, cartItem.ItemID, cartItem.Quantity)
	return
}

func (service *cartServiceStorer) UpdateQuantity(ctx context.Context, cartID string, quantity int) (cart dto.Cart, err error) {
	accountId := ctx.Value(dto.AccountID).(string)
	err = service.db.GetContext(ctx, &cart, updateCartQuantityQuery, cartID, quantity, accountId)
	return
}

func (service *cartServiceStorer) ListCartItems(ctx context.Context, accountID string, pageNumber int) (cartItems []dto.CartItems, err error) {
	offset := (pageNumber - 1) * service.itemPerPage
	err = service.db.SelectContext(ctx, &cartItems, listCartItemsQuery, accountID, offset, service.itemPerPage)
	return
}
