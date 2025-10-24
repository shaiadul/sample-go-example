package handlers

import (
	"encoding/json"
	"net/http"
	"sample-go/database"
	"sample-go/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var NewProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&NewProduct)
	if err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}

	NewProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, NewProduct)
	util.SentData(w, NewProduct, http.StatusCreated)

}
