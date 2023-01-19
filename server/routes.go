package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/megre/accounts"
	"github.com/megre/app"
	"github.com/megre/cart"
	"github.com/megre/dto"
	"github.com/megre/items"
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

	// accounts Router
	accountsRouter := r.PathPrefix("/accounts").Subrouter().StrictSlash(false)
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

	// CART routers
	cartRouter := r.PathPrefix("/cart").Subrouter().StrictSlash(false)
	cartRouter.
		Handle("", Adapt(http.HandlerFunc(cart.HandleUpsertItemToCart(dependencies.CartService)),
			middleware.CheckAllowedRole(dto.ResourceIdentifier{Module: "cart", Resource: "upsert"}),
			middleware.AuthCheck(app.GetConfiguration().JWT_SECRET))).
		Methods("POST")

	cartRouter.
		Handle("", Adapt(http.HandlerFunc(cart.HandleListCartItem(dependencies.CartService)),
			middleware.CheckAllowedRole(dto.ResourceIdentifier{Module: "cart", Resource: "get"}),
			middleware.AuthCheck(app.GetConfiguration().JWT_SECRET))).
		Methods("GET")

	itemsRouter := r.PathPrefix("/item").Subrouter().StrictSlash(false)
	itemsRouter.
		Handle("", Adapt(http.HandlerFunc(items.HandleAddItem(dependencies.ItemService)),
			middleware.CheckAllowedRole(dto.ResourceIdentifier{Module: "items", Resource: "add"}),
			middleware.AuthCheck(app.GetConfiguration().JWT_SECRET))).
		Methods("POST")
}

func GetRouter() *mux.Router {
	return r
}
