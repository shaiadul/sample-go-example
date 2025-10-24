package handlers

import (
	"net/http"
	"sample-go/product"
	"sample-go/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	util.SentData(w, product.ProductList, http.StatusOK)
}
