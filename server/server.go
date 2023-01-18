package server

import (
	"fmt"
	"net/http"

	"github.com/megre/app"
)

func StartHTTPServer() {
	app.InitaliseApp()
	initialiseRoute()

	configuration := app.GetConfiguration()
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", configuration.APP_HOST, configuration.APP_PORT),
		Handler: GetRouter(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("error starting server, error: %s", err.Error()))
	}

	fmt.Println("successfully connected to the database")
}
