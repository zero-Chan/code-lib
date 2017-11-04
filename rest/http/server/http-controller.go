package server

import (
	"net/http"

	reflect_util "code-lib/reflect-util"
	"code-lib/rest"
)

type HTTPController struct {
	svr     *HTTPServer
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
	var (
		restResp *rest.RestResponse
	)
	respw.WriteHeader(http.StatusOK)

	// new channel
	httpch := NewHTTPChannel(respw, req)

	// new handler
	execHandler, ok := reflect_util.NewInterface(this.Handler).(HTTPHandler)
	if !ok {
		respw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// exec handler
	execHandler.InitFromHTTP(httpch)

	for {
		restResp = execHandler.Prepare()
		if !restResp.IsOk() {
			break
		}

		restResp = execHandler.Handle()
		if !restResp.IsOk() {
			break
		}

		restResp = execHandler.Finish()
		if !restResp.IsOk() {
			break
		}

		break
	}

	jsonbytes, _ := restResp.Marshal2JSON()
	respw.Write(jsonbytes)
}
