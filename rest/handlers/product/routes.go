package product

import (
	"net/http"
	"sample-go/rest/middleware"
)

func (h *Handler) ProductRoutes(router *http.ServeMux, manager *middleware.Manager) {

	router.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
		),
	)

	router.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			h.middlewares.AuthenticateJWT,
		),
	)

	router.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProduct),
		),
	)

	router.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
			h.middlewares.AuthenticateJWT,
		),
	)

	router.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			h.middlewares.AuthenticateJWT,
		),
	)

}
