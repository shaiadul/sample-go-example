package cmd

import (
	"fmt"
	"net/http"
	"sample-go/global_router"
	"sample-go/handlers"
)

func Serve() {
	router := http.NewServeMux()

	router.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	router.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	router.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductById))

	globalRouter := global_router.GlobalRouter(router)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error Starting Server", err)
	}
}
