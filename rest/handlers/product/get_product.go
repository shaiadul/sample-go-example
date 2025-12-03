package product

import (
	"net/http"
	"sample-go/database"
	"sample-go/util"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")

	pid, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	product := database.Get(pid)
	if product == nil {
		util.SentError(w, "Product not found", http.StatusNotFound)
		return
	}

	util.SentData(w, product, http.StatusOK)

}
