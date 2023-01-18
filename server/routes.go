package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/megre/accounts"
	"github.com/megre/app"
	"github.com/megre/middleware"
)

var r *mux.Router

type Adapters func(http.Handler) http.Handler

func Adapt(handler http.Handler, adapters ...Adapters) http.Handler {
	for i := len(adapters); i > 0; i-- {
		handler = adapters[i-1](handler)
	}

	return handler
}

func initialiseRoute() {
	dependencies := app.GetDependencies()
	r = mux.NewRouter()

	accountsRouter := r.PathPrefix("/accounts").Subrouter()

	accountsRouter.
		HandleFunc("/create", Adapt(http.HandlerFunc(accounts.HandleCreateAccount(dependencies.AccountService))).ServeHTTP).
		Methods("POST")

	accountsRouter.
		HandleFunc("/login", Adapt(http.HandlerFunc(accounts.HandleUserLogin(dependencies.AccountService))).ServeHTTP).
		Methods("POST")

	accountsRouter.HandleFunc("/suspend/{id}",
		Adapt(
			http.HandlerFunc(accounts.HandleSuspendAccount(dependencies.AccountService)),
			middleware.AuthCheck(app.GetConfiguration().JWT_SECRET),
		).ServeHTTP).Methods("POST")
}

func GetRouter() *mux.Router {
	return r
}
