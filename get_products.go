package main

import "net/http"

func getProducts(w http.ResponseWriter, r *http.Request) {

	sentData(w, productList, http.StatusOK)
}
