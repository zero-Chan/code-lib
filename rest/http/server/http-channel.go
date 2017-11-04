package server

import (
	"net/http"

	"github.com/pborman/uuid"
)

type HTTPChannel struct {
	ID        string
	request   *http.Request
	responsew http.ResponseWriter
}

func CreateHTTPChannel(respw http.ResponseWriter, req *http.Request) HTTPChannel {
	ch := HTTPChannel{
		ID:        uuid.New(),
		request:   req,
		responsew: respw,
	}
	return ch
}

func NewHTTPChannel(respw http.ResponseWriter, req *http.Request) *HTTPChannel {
	ch := CreateHTTPChannel(respw, req)
	return &ch
}
