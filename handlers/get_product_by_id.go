package handlers

import (
	"net/http"
	"sample-go/database"
	"sample-go/util"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")

	pid, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pid {
			util.SentData(w, product, http.StatusOK)
			return
		}

	}

	util.SentData(w, "Product not found", http.StatusNotFound)

}
