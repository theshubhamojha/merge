package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/megre/dto"
	"github.com/megre/merrors"
)

func HandleCreateAccount(accountService AccountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dto.CreateAccountRequest
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

		serviceError := request.ValidateCreateAccountRequest()
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

		response, serviceError := accountService.CreateAccount(r.Context(), request)
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
				Message: "account created successfully",
			},
			http.StatusOK,
		)
	}
}

func HandleSuspendAccount(accountService AccountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		request, ok := vars["id"]
		if !ok {
			dto.SendAPIResponse(w,
				dto.APIResponse{
					Message:   "account id not specified in request url",
					ErrorCode: merrors.BadRequets,
				},
				http.StatusBadRequest,
			)
			return
		}

		serviceError := accountService.SuspendAccount(r.Context(), request)
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
				Message: "account suspended successfully",
			},
			http.StatusOK,
		)
	}
}

func HandleUserLogin(accountService AccountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dto.LoginRequest
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

		response, responseErr := accountService.LoginAccount(r.Context(), request)
		if responseErr != nil {
			if responseErr.ErrorCode == merrors.IncorrectCredentials {
				dto.SendAPIResponse(w,
					dto.APIResponse{
						Message:   "invalid user credentials",
						ErrorCode: merrors.IncorrectCredentials,
					},
					http.StatusUnauthorized,
				)
				return
			}

			dto.SendAPIResponse(w,
				dto.APIResponse{
					Message:   "internal server error",
					ErrorCode: merrors.InternalServerError,
				},
				http.StatusInternalServerError,
			)
		}

		dto.SendAPIResponse(w,
			dto.APIResponse{
				Data:    response,
				Message: "account login successful",
			},
			http.StatusOK,
		)
	}
}
