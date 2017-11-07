package server

import (
	"net/http"

	reflect_util "code-lib/reflect-util"
	"code-lib/rest"
)

type HTTPController struct {
	svr     *HTTPServer
	handler HTTPHandler
}

func createHTTPController(svr *HTTPServer, hdl HTTPHandler) HTTPController {
	hc := HTTPController{
		svr:     svr,
		handler: hdl,
	}
	return hc
}

func newHTTPController(svr *HTTPServer, hdl HTTPHandler) *HTTPController {
	hc := createHTTPController(svr, hdl)
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
	execHandler, ok := reflect_util.NewInterface(this.handler).(HTTPHandler)
	if !ok {
		respw.WriteHeader(http.StatusInternalServerError)
		respw.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	for {
		// exec handler
		restResp = execHandler.InitFromHTTP(httpch)
		if !restResp.IsOk() {
			break
		}

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
