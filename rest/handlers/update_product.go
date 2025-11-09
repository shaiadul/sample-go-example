package handlers

import (
	"encoding/json"
	"net/http"
	"sample-go/database"
	"sample-go/util"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pid, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	var NewProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&NewProduct)
	if err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}

	NewProduct.ID = pid

	database.Update(NewProduct)

	util.SentData(w, "Successfully updated product", http.StatusOK)

}
