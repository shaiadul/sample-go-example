package handlers

import (
	"net/http"
	"sample-go/database"
	"sample-go/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	util.SentData(w, database.ProductList, http.StatusOK)
}
