package cart

import (
	"context"
	"database/sql"
	"math"

	"github.com/megre/dto"
	"github.com/megre/items"
	"github.com/megre/merrors"
	"github.com/megre/store"
	"github.com/megre/utils"
)

type CartService interface {
	UpsertCartItem(ctx context.Context, itemId string, accountID string, quantity int) (cartItem dto.Cart, err *merrors.Error)
	ListCartItems(ctx context.Context, accountID string, pageNumber int) (cartItem []dto.CartItems, err *merrors.Error)
}

type cartService struct {
	itemService items.ItemService
	cartStore   store.CartServiceStore
}

func NewCartService(
	itemService items.ItemService,
	cartStore store.CartServiceStore,
) CartService {
	return &cartService{
		itemService: itemService,
		cartStore:   cartStore,
	}
}

func (service *cartService) UpsertCartItem(ctx context.Context, itemId string, accountId string, quantity int) (cartItem dto.Cart, err *merrors.Error) {
	item, getItemError := service.itemService.GetItem(ctx, itemId)
	if getItemError != nil {
		if getItemError.ErrorCode == merrors.DataNotExist {
			err = &merrors.Error{
				ErrorCode: merrors.DataNotExist,
				Message:   "item doesn't exist",
			}
			return
		}

		err = &merrors.Error{
			ErrorCode: merrors.InternalServerError,
			Message:   "unable to get items",
		}
		return
	}

	if item.Quantity < quantity {
		err = &merrors.Error{
			Message:   "requested quantity out of stock, please lower the quantity",
			ErrorCode: merrors.OutOfStock,
		}
		return
	}

	// checking if the cart already exisit
	cart, dbError := service.cartStore.GetCartDetail(ctx, itemId, accountId)
	if dbError != nil && dbError == sql.ErrNoRows {
		cart, insertErr := service.insertItemToCart(ctx, item.ID, accountId, quantity)
		if insertErr != nil {
			err = &merrors.Error{
				ErrorCode: merrors.InternalServerError,
				Message:   "error upserting item to cart",
				Error:     insertErr,
			}
			return
		}
		return cart, nil
	}
	if dbError != nil {
		err = &merrors.Error{
			ErrorCode: merrors.InternalServerError,
			Message:   "error upserting item to cart",
		}
		return
	}

	if cart.Quantity+quantity > item.Quantity {
		err = &merrors.Error{
			Message:   "requested quantity out of stock, please lower the quantity",
			ErrorCode: merrors.OutOfStock,
		}
		return
	}

	resultantQuantity := int(math.Max((float64(cart.Quantity + quantity)), 0))

	cart, dbError = service.updateCartQuantity(ctx, resultantQuantity, cart.ID)
	if dbError != nil {
		err = &merrors.Error{
			ErrorCode: merrors.InternalServerError,
			Message:   "error upserting item to cart",
		}
		return
	}

	return cart, nil
}

func (service *cartService) ListCartItems(ctx context.Context, accountID string, pageNumber int) (cartItems []dto.CartItems, err *merrors.Error) {
	cartItems, dbError := service.cartStore.ListCartItems(ctx, accountID, pageNumber)
	if dbError != nil {
		// if no data is found, return empty array and no error
		if dbError == sql.ErrNoRows {
			return cartItems, nil
		}

		err = &merrors.Error{
			Message:   "error fetching cart item data",
			ErrorCode: merrors.InternalServerError,
		}
		return
	}

	return
}

func (service *cartService) insertItemToCart(ctx context.Context, itemID string, accountID string, quantity int) (cart dto.Cart, err error) {
	cartItem := dto.Cart{
		ID:        utils.GenerateRandomID(ctx, "cart_"),
		AccountID: accountID,
		ItemID:    itemID,
		Quantity:  quantity,
	}

	cart, err = service.cartStore.InsertCart(ctx, cartItem)
	return
}

func (service *cartService) updateCartQuantity(ctx context.Context, resultantQuantity int, cartID string) (cart dto.Cart, err error) {
	cart, err = service.cartStore.UpdateQuantity(ctx, cartID, resultantQuantity)
	return
}
