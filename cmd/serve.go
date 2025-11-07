package cmd

import (
	"fmt"
	"net/http"
	"sample-go/middleware"
)

func Serve() {

	router := http.NewServeMux()

	manager := middleware.NewManager()

	wrapperRouter := manager.WrapRouter(
		router,
		middleware.Logger,
		middleware.PreFlight,
		middleware.Cors,
	)

	initRoutes(router, manager)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", wrapperRouter)
	if err != nil {
		fmt.Println("Error Starting Server", err)
	}
}
