package middleware

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/megre/dto"
	"github.com/megre/merrors"
	"github.com/megre/utils"
)

func AuthCheck(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			isValid, claims := isAuthHeaderValid(authHeader, jwtSecret)
			if !isValid {
				dto.SendAPIResponse(w,
					dto.APIResponse{
						Message:   "unauthorized access",
						ErrorCode: merrors.Unauthorized,
					},
					http.StatusUnauthorized,
				)
				return
			}

			ctx := context.WithValue(r.Context(), role, claims["role"])
			ctx = context.WithValue(ctx, email, claims["email"])
			ctx = context.WithValue(ctx, accountID, claims["account_id"])

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func isAuthHeaderValid(token string, jwtSecret string) (isValid bool, claims jwt.MapClaims) {
	claims, err := utils.VerifyJWTToken(token, jwtSecret)
	if err != nil {
		isValid = false
		return
	}

	err = claims.Valid()
	if err != nil {
		isValid = false
		return
	}

	return true, claims
}
