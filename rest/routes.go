package rest

import (
	"net/http"
	"sample-go/rest/handlers"
	"sample-go/rest/middleware"
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
			middleware.AuthenticateJWT,
		),
	)

	router.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
		),
	)

	router.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProduct),
			middleware.AuthenticateJWT,
		),
	)

	router.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProduct),
			middleware.AuthenticateJWT,
		),
	)

	// user Routes here
	router.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(handlers.CreateUser),
		),
	)

	router.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		),
	)
}
