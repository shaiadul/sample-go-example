package product

import (
	"net/http"
	"sample-go/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.productRepo.List()
	if err != nil {
		util.SentError(w, http.StatusInternalServerError, "Error while fetching products")
		return
	}
	util.SentData(w, http.StatusOK, productList)
}
