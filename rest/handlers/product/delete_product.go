package product

import (
	"net/http"
	"sample-go/util"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)
	if err != nil {
		util.SentError(w, http.StatusBadRequest, "Invalid product id")
		return
	}

	err = h.productRepo.Delete(pId)
	if err != nil {

		util.SentError(w, http.StatusInternalServerError, "Failed to delete the product")

		return
	}

	util.SentData(w, http.StatusOK, "Successfully deleted the product")

}
