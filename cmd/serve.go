package cmd

import (
	"sample-go/config"
	"sample-go/rest"
)

func Serve() {

	cnf := config.GetConfig()

	rest.Start(cnf)
}
