package resthttp

import (
	"net/http"

	"code-lib/gerror"
)

type HTTPMux interface {
	FindBuilder(req *http.Request) (builder HTTPBuilder, gerr *gerror.GError)
}
