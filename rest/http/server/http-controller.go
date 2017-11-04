package server

import (
	"net/http"

	reflect_util "code-lib/reflect-util"
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
	// new global processor
	var execProcessor HTTPProcessor
	if this.svr.processor != nil {
		var ok bool
		execProcessor, ok = reflect_util.NewInterface(this.svr.processor).(HTTPProcessor)
		if !ok {
			respw.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// new handler
	execHandler, ok := reflect_util.NewInterface(this.Handler).(HTTPHandler)
	if !ok {
		respw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// new channel
	httpch := NewHTTPChannel(respw, req)

	// exec
	if execProcessor != nil {
		execProcessor.InitFromHTTP(httpch)

		execProcessor.Prepare()
	}

	execHandler.InitFromHTTP(httpch)

	execHandler.Prepare()
	execHandler.Handle()
	execHandler.Finish()

	if execProcessor != nil {
		execProcessor.Finish()
	}
	respw.WriteHeader(http.StatusOK)
}
