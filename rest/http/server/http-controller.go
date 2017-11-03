package server

import (
	"net/http"
)

type HTTPController struct {
	svr     *HTTPMuxServer
	Handler HTTPHandler
}

func CreateHTTPController() HTTPController {
	hc := HTTPController{}
	return hc
}

func NewHTTPController() *HTTPController {
	hc := CreateHTTPController()
	return &hc
}

func (this *HTTPController) ServeHTTP(respw http.ResponseWriter, req *http.Request) {

}
