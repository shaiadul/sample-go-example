package cmd

import (
	"net/http"
	"sample-go/handlers"
	"sample-go/middleware"
)

func initRoutes(router *http.ServeMux, manager *middleware.Manager) {
	router.Handle(
		"GET /test",
		manager.With(
			http.HandlerFunc(handlers.Test),
		),
	)

	router.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)

	router.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)

	router.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductById),
		),
	)
}
