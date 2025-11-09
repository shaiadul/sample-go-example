package rest

import (
	"fmt"
	"net/http"
	"os"
	"sample-go/config"
	"sample-go/rest/middleware"
	"strconv"
)

func Start(cnf config.Config) {

	router := http.NewServeMux()

	manager := middleware.NewManager()

	wrapperRouter := manager.WrapRouter(
		router,
		middleware.Logger,
		middleware.PreFlight,
		middleware.Cors,
	)

	initRoutes(router, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server is running on port:", addr)
	err := http.ListenAndServe(addr, wrapperRouter)
	if err != nil {
		fmt.Println("Error Starting Server", err)
		os.Exit(1)
	}
}
