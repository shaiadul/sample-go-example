package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {

	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {

	n := next

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	// for _, gloabalMiddleware := range mngr.globalMiddlewares {
	// 	n = gloabalMiddleware(n)
	// }

	return n

}

func (mngr *Manager) WrapRouter(next http.Handler, middlewares ...Middleware) http.Handler {

	n := next

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	return n

}
