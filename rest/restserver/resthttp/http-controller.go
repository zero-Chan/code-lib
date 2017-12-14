package resthttp

import (
	"net/http"

	"code-lib/rest"
)

type HTTPController struct {
	svr     *HTTPServer
	builder HTTPBuilder
}

func createHTTPController(svr *HTTPServer, builder HTTPBuilder) HTTPController {
	hc := HTTPController{
		svr:     svr,
		builder: builder,
	}
	return hc
}

func newHTTPController(svr *HTTPServer, builder HTTPBuilder) *HTTPController {
	hc := createHTTPController(svr, builder)
	return &hc
}

func (this *HTTPController) ServeHTTP(respw http.ResponseWriter, req *http.Request) {
	var (
		restResp *rest.RestResponse
		executor rest.Executor
	)
	respw.WriteHeader(http.StatusOK)

	// new channel
	httpch := NewHTTPChannel(respw, req)

	for {
		// exec handler
		executor, restResp = this.builder.BuildExectorFromHTTP(httpch)
		if !restResp.IsOk() {
			break
		}

		restResp = executor.Prepare()
		if !restResp.IsOk() {
			break
		}

		restResp = executor.Exec()
		if !restResp.IsOk() {
			break
		}

		restResp = executor.Finish()
		if !restResp.IsOk() {
			break
		}

		break
	}

	bytes, _ := restResp.Marshal()
	respw.Write(bytes)
}
