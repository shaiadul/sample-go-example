package rest

import (
	"fmt"
	"net/http"
	"os"
	"sample-go/config"
	"sample-go/rest/handlers/product"
	"sample-go/rest/handlers/user"
	"sample-go/rest/middleware"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewSeever(
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {

	router := http.NewServeMux()

	manager := middleware.NewManager()

	wrapperRouter := manager.WrapRouter(
		router,
		middleware.Logger,
		middleware.PreFlight,
		middleware.Cors,
	)

	// Define routes
	server.productHandler.ProductRoutes(router, manager)
	server.userHandler.RegisterRoutes(router, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server is running on port:", addr)
	err := http.ListenAndServe(addr, wrapperRouter)
	if err != nil {
		fmt.Println("Error Starting Server", err)
		os.Exit(1)
	}
}
