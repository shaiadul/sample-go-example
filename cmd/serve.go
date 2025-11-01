package cmd

import (
	"fmt"
	"net/http"
	"sample-go/global_router"
	"sample-go/middleware"
)

func Serve() {

	manager := middleware.NewManager()

	manager.Use(middleware.Logger, middleware.Hudai)

	router := http.NewServeMux()

	initRoutes(router, manager)

	globalRouter := global_router.GlobalRouter(router)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error Starting Server", err)
	}
}
