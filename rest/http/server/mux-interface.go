package server

import (
	"net/http"

	"code-lib/gerror"
)

type HTTPMux interface {
	FindHandler(req *http.Request) (hdl HTTPHandler, gerr *gerror.GError)
}
