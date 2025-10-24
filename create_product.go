package main

import (
	"encoding/json"
	"net/http"
)

func createProduct(w http.ResponseWriter, r *http.Request) {

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
