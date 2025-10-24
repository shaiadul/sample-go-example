package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Our Website | api")
}

type Product struct {
	ID          int     `json:"id"` // if i use id is make privete then it will give error , if need small case then use tag
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {

	handleCors(w)
	handlePreflightRequest(w, r)

	if r.Method != "GET" {
		http.Error(w, "Please Give me Valid Method", http.StatusMethodNotAllowed)
		return
	}
	// encode via encoder package to json format
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func createProduct(w http.ResponseWriter, r *http.Request) {

	handleCors(w)
	handlePreflightRequest(w, r)

	if r.Method != "POST" {
		http.Error(w, "Please Give me POST Method", http.StatusMethodNotAllowed)
		return
	}

	var NewProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&NewProduct)
	if err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}

	NewProduct.ID = len(productList) + 1
	productList = append(productList, NewProduct)
	sentData(w, NewProduct, http.StatusCreated)

}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-control-Allow-Headers", "Content-Type, custom-header") // if i have custom header then add it here , other wise remove it only use Content-Type
	w.Header().Set("Content-Type", "application/json")
}

func handlePreflightRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusCreated)
	}
}

func sentData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {

	router := http.NewServeMux()
	router.Handle("GET /products", http.HandlerFunc(getProducts))
	router.Handle("POST /products/create", http.HandlerFunc(createProduct))
	router.Handle("GET /about", http.HandlerFunc(handleAbout))
	router.Handle("GET /hello", http.HandlerFunc(handleHello))

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", router)
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
