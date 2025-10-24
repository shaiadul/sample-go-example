package handlers

import (
	"encoding/json"
	"net/http"
	"sample-go/product"
	"sample-go/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var NewProduct product.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&NewProduct)
	if err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}

	NewProduct.ID = len(product.ProductList) + 1
	product.ProductList = append(product.ProductList, NewProduct)
	util.SentData(w, NewProduct, http.StatusCreated)

}
