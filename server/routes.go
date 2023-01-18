package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/megre/accounts"
	"github.com/megre/app"
	"github.com/megre/dto"
	"github.com/megre/middleware"
)

var r *mux.Router

type Adapters func(http.Handler) http.Handler

func Adapt(handler http.Handler, adapters ...Adapters) http.Handler {
	for _, adapter := range adapters {
		handler = adapter(handler)
	}

	return handler
}

func initialiseRoute() {
	dependencies := app.GetDependencies()
	r = mux.NewRouter()

	accountsRouter := r.PathPrefix("/accounts").Subrouter()

	accountsRouter.
		Handle("/create", Adapt(http.HandlerFunc(accounts.HandleCreateAccount(dependencies.AccountService)))).
		Methods("POST")

	accountsRouter.
		Handle("/login", Adapt(http.HandlerFunc(accounts.HandleUserLogin(dependencies.AccountService)))).
		Methods("POST")

	accountsRouter.Handle("/suspend/{id}",
		Adapt(
			http.HandlerFunc(accounts.HandleSuspendAccount(dependencies.AccountService)),
			middleware.CheckAllowedRole(dto.ResourceIdentifier{Module: "accounts", Resource: "suspend"}),
			middleware.AuthCheck(app.GetConfiguration().JWT_SECRET),
		)).
		Methods("POST")
}

func GetRouter() *mux.Router {
	return r
}
