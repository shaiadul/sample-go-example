package middleware

import "net/http"

func Cors(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-control-Allow-Headers", "Content-Type, custom-header") // if i have custom header, then add it here , other wise remove it only use Content-Type
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)

	})
}
