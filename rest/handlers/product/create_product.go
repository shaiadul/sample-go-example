package product

import (
	"encoding/json"
	"net/http"
	"sample-go/repo"
	"sample-go/util"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {

		util.SentError(w, http.StatusBadRequest, "Invalid JSON Data")
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		util.SentError(w, http.StatusInternalServerError, "Error while creating product")
		return
	}

	util.SentData(w, http.StatusOK, createdProduct)

}
