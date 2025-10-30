package cmd

import (
	"fmt"
	"net/http"
	"sample-go/global_router"
	"sample-go/handlers"
	"sample-go/middleware"
)

func Serve() {
	router := http.NewServeMux()

	router.Handle("GET /test", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.Test))))

	router.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	router.Handle("POST /products", middleware.Logger(http.HandlerFunc(handlers.CreateProduct)))
	router.Handle("GET /products/{id}", middleware.Logger(http.HandlerFunc(handlers.GetProductById)))

	globalRouter := global_router.GlobalRouter(router)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error Starting Server", err)
	}
}
