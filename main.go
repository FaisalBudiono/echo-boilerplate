package main

import (
	"echo-boilerplate/bootstraps"
)

func main() {
	bootstraps.BindEnv()

	server := bootstraps.NewWebServer()
	server.ListenAndServe()
}
