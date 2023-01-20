package items

import (
	"encoding/json"
	"net/http"

	"github.com/megre/dto"
	"github.com/megre/merrors"
)

func HandleAddItem(itemService ItemService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dto.ItemRequest
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

		serviceError := request.ValidateRequest()
		if serviceError != nil {
			dto.SendAPIResponse(w,
				dto.APIResponse{
					Message:   serviceError.Message,
					ErrorCode: serviceError.ErrorCode,
				},
				http.StatusBadRequest,
			)
			return
		}

		response, serviceError := itemService.AddItem(r.Context(), request)
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
				Data:    response,
				Message: "item added successfully",
			},
			http.StatusOK,
		)
	}
}
