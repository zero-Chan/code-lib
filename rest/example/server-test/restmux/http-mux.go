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

	//	HTTPMux.MPHTTPMux.RegisterHandler(
	//		handlers.NewClientDel(),
	//		mux.MPHTTPMuxRule{
	//			Path:   "/client/del",
	//			Method: http.MethodPost,
	//		},
	//	)
}
