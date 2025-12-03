package user

import (
	"net/http"
	"sample-go/rest/middleware"
)

func (h *Handler) RegisterRoutes(router *http.ServeMux, manager *middleware.Manager) {

	router.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(h.CreateUser),
		),
	)

	router.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(h.Login),
		),
	)
}
