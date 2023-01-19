package cart

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/megre/dto"
	"github.com/megre/merrors"
)

func HandleUpsertItemToCart(cartService CartService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dto.UpsertCartRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			dto.SendAPIResponse(w,
				dto.APIResponse{
					Message:   "error reading request body",
					ErrorCode: merrors.BadRequets,
				},
				http.StatusBadRequest,
			)
			return
		}

		accountID := r.Context().Value(dto.AccountID).(string)
		quantity := request.Quantity
		itemID := request.ItemId

		cartItem, serviceError := cartService.UpsertCartItem(r.Context(), itemID, accountID, quantity)
		if serviceError != nil {
			dto.SendAPIResponse(w,
				dto.APIResponse{
					Message:   serviceError.Message,
					ErrorCode: serviceError.ErrorCode,
				},
				http.StatusInternalServerError,
			)
			return
		}

		dto.SendAPIResponse(w,
			dto.APIResponse{
				Data:    cartItem,
				Message: "the operation was successful",
			},
			http.StatusOK,
		)
	}
}

func HandleListCartItem(cartService CartService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		accountID := r.Context().Value(dto.AccountID).(string)
		page := r.URL.Query().Get("page")
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			dto.SendAPIResponse(w,
				dto.APIResponse{
					Message:   "invalid `page` query param",
					ErrorCode: merrors.BadRequets,
				},
				http.StatusBadRequest,
			)
		}

		cartItems, serviceError := cartService.ListCartItems(r.Context(), accountID, pageInt)
		if serviceError != nil {
			dto.SendAPIResponse(w,
				dto.APIResponse{
					Message:   serviceError.Message,
					ErrorCode: serviceError.ErrorCode,
				},
				http.StatusInternalServerError,
			)
			return
		}

		dto.SendAPIResponse(w,
			dto.APIResponse{
				Data: cartItems,
			},
			http.StatusOK,
		)
	}
}
