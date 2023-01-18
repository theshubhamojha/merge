package middleware

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/megre/app"
	"github.com/megre/dto"
)

var db *sqlx.DB = app.GetDB()

func CheckAllowedRole(resource dto.ResourceIdentifier) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL)
			next.ServeHTTP(w, r)
		})
	}
}
