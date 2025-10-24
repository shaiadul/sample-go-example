package main

import (
	"fmt"
	"net/http"
	"sample-go/global_router"
	"sample-go/handlers"
	"sample-go/product"
)

func main() {

	router := http.NewServeMux()

	router.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	router.Handle("POST /products/create", http.HandlerFunc(handlers.CreateProduct))

	globalRouter := global_router.GlobalRouter(router)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error Starting Server", err)
	}
}

func init() {
	pro01 := product.Product{
		ID:          1,
		Title:       "Product 1",
		Description: "This is product 1 description",
		Price:       100.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro02 := product.Product{
		ID:          2,
		Title:       "Product 2",
		Description: "This is product 2 description",
		Price:       200.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro03 := product.Product{
		ID:          3,
		Title:       "Product 3",
		Description: "This is product 3 description",
		Price:       300.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro04 := product.Product{
		ID:          4,
		Title:       "Product 4",
		Description: "This is product 4 description",
		Price:       400.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro05 := product.Product{
		ID:          5,
		Title:       "Product 5",
		Description: "This is product 5 description",
		Price:       500.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	product.ProductList = append(product.ProductList, pro01, pro02, pro03, pro04, pro05)
}
