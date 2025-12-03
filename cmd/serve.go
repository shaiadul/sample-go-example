package cmd

import (
	"sample-go/config"
	"sample-go/rest"
	"sample-go/rest/handlers/product"
	"sample-go/rest/handlers/user"
	"sample-go/rest/middleware"
)

func Serve() {

	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()

	srver := rest.NewSeever(cnf, productHandler, userHandler)

	srver.Start()
}
