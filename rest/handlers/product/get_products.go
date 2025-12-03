package product

import (
	"net/http"
	"sample-go/database"
	"sample-go/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	util.SentData(w, database.List(), http.StatusOK)
}
