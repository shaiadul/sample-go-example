package product

import (
	"encoding/json"
	"net/http"
	"sample-go/repo"
	"sample-go/util"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)
	if err != nil {
		util.SentError(w, http.StatusBadRequest, "Invalid product id")

		return
	}

	var req ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		util.SentError(w, http.StatusBadRequest, "Invalid JSON Data")
		return
	}

	_, err = h.productRepo.Update(repo.Product{
		ID:          pId,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		util.SentError(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	util.SentData(w, http.StatusOK, "Successfully updated product")

}
