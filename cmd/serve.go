package cmd

import (
	"sample-go/config"
	"sample-go/repo"
	"sample-go/rest"
	"sample-go/rest/handlers/product"
	"sample-go/rest/handlers/user"
	"sample-go/rest/middleware"
)

func Serve() {

	cnf := config.GetConfig()

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	srver := rest.NewSeever(cnf, productHandler, userHandler)

	srver.Start()
}
