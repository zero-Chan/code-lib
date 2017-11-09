package restmux

import (
	"net/http"

	"code-lib/rest/restserver/resthttp/mux"

	"rest_example/server-test/executor/client-exec"
)

var (
	HTTPMux *httpMux
)

type httpMux struct {
	*mux.MPHTTPMux
}

func init() {
	HTTPMux = &httpMux{
		MPHTTPMux: mux.NewHTTPPathMux(),
	}

	regist()
}

func regist() {
	HTTPMux.MPHTTPMux.RegisterBuilder(
		client_exec.NewClientAddBuilder(),
		mux.MPHTTPMuxRule{
			Path:   "/client/add",
			Method: http.MethodPost,
		},
	)

	HTTPMux.MPHTTPMux.RegisterBuilder(
		client_exec.NewClientGetBuilder(),
		mux.MPHTTPMuxRule{
			Path:   "/client/get",
			Method: http.MethodPost,
		},
	)
}
