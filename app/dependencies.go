package app

import (
	"github.com/megre/accounts"
	"github.com/megre/cart"
	"github.com/megre/items"
	"github.com/megre/store"
)

type Dependencies struct {
	AccountService accounts.AccountService
	CartService    cart.CartService
	ItemService    items.ItemService
}

func initialiseDependencies() *Dependencies {
	db := GetDB()
	configuration = GetConfiguration()

	accountStorerService := store.NewAccountStorerService(db)
	accountService := accounts.NewAccountService(configuration.JWT_SECRET, configuration.JWT_EXPIRY, accountStorerService)

	itemServiceStorer := store.NewItemServiceStore(db)
	itemService := items.NewItemService(itemServiceStorer)

	cartServiceStorer := store.NewCartServiceStore(db, configuration.DATA_PER_PAGE)
	cartService := cart.NewCartService(itemService, cartServiceStorer)

	return &Dependencies{
		AccountService: accountService,
		CartService:    cartService,
		ItemService:    itemService,
	}
}
