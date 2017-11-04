package server

import (
	"net/http"

	"github.com/pborman/uuid"
)

type HTTPChannel struct {
	ID        string
	Request   *http.Request
	Responsew http.ResponseWriter
}

func CreateHTTPChannel(respw http.ResponseWriter, req *http.Request) HTTPChannel {
	ch := HTTPChannel{
		ID:        uuid.New(),
		Request:   req,
		Responsew: respw,
	}
	return ch
}

func NewHTTPChannel(respw http.ResponseWriter, req *http.Request) *HTTPChannel {
	ch := CreateHTTPChannel(respw, req)
	return &ch
}
