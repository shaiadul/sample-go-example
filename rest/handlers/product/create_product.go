package product

import (
	"encoding/json"
	"net/http"
	"sample-go/database"
	"sample-go/util"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var NewProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&NewProduct)
	if err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}

	createdProduct := database.Store(NewProduct)

	util.SentData(w, createdProduct, http.StatusOK)

}
