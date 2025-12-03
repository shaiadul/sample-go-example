package product

import (
	"net/http"
	"sample-go/database"
	"sample-go/util"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pid, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	database.Delete(pid)

	util.SentData(w, "Successfully deleted the product", http.StatusOK)

}
