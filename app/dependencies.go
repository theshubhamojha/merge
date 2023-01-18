package app

import (
	"github.com/megre/accounts"
	"github.com/megre/store"
)

type Dependencies struct {
	AccountService accounts.AccountService
}

func initialiseDependencies() *Dependencies {
	db := GetDB()
	configuration = GetConfiguration()

	accountStorerService := store.NewAccountStorerService(db)
	accountService := accounts.NewAccountService(configuration.JWT_SECRET, configuration.JWT_EXPIRY, accountStorerService)

	return &Dependencies{
		AccountService: accountService,
	}
}
