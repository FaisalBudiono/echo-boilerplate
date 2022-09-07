package bootstraps

import (
	"net/http"
)

func NewWebServer() http.Server {
	return http.Server{
		Addr:         getAddr(),
		Handler:      createRoute(),
		ReadTimeout:  getHttpTimeout(),
		WriteTimeout: getHttpTimeout(),
	}
}
