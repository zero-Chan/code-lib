package server

import (
	"net/http"
)

type HTTPMux interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}
