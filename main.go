package main

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Our Website | api")
}

func main() {

	route := http.NewServeMux()
	route.HandleFunc("/", handleHello)
	route.HandleFunc("/about", handleAbout)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", route)
	if err != nil {
		fmt.Println("Error Starting Server", err)
	}
}
