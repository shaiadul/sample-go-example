package product

import (
	"net/http"
	"sample-go/util"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)
	if err != nil {
		util.SentError(w, http.StatusBadRequest, "Invalid product id")
		return
	}

	product, err := h.productRepo.Get(pId)
	if err != nil {
		util.SentError(w, http.StatusInternalServerError, "Error getting product")
		return
	}
	if product == nil {
		util.SentError(w, http.StatusNotFound, "Product not found")
		return
	}

	util.SentData(w, http.StatusOK, product)

}
