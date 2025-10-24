package main

import (
	"fmt"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.Handle("GET /products", http.HandlerFunc(getProducts))
	router.Handle("POST /products/create", http.HandlerFunc(createProduct))

	globalRouter := globalRouter(router)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error Starting Server", err)
	}
}

func init() {
	pro01 := Product{
		ID:          1,
		Title:       "Product 1",
		Description: "This is product 1 description",
		Price:       100.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro02 := Product{
		ID:          2,
		Title:       "Product 2",
		Description: "This is product 2 description",
		Price:       200.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro03 := Product{
		ID:          3,
		Title:       "Product 3",
		Description: "This is product 3 description",
		Price:       300.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro04 := Product{
		ID:          4,
		Title:       "Product 4",
		Description: "This is product 4 description",
		Price:       400.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro05 := Product{
		ID:          5,
		Title:       "Product 5",
		Description: "This is product 5 description",
		Price:       500.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	productList = append(productList, pro01, pro02, pro03, pro04, pro05)
}
