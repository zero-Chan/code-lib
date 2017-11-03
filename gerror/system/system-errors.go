package system_err

import (
	"fmt"
	"net/http"

	"code-lib/gerror"
)

// SystemErr 以 10000　为基准

const (
	ErrSystemCode = 10000

	// HTTPErr
	ErrHTTPCode           = 20100
	ErrHTTPRespStatusCode = 20101
)

// =====
// EncodingErr:
// Code: range [20000, 30000)
func ErrSystem(reason error) (gerr *gerror.GError) {
	gerr = gerror.NewGError()
	gerr.Code = ErrSystemCode
	gerr.Label = "ErrSystem"
	gerr.Message = reason.Error()
	return
}

// ==========
// JSONErr:
// Code: range [20100, 20199)
func ErrHTTP(reason error) (gerr *gerror.GError) {
	gerr = gerror.NewGError()
	gerr.Code = ErrHTTPCode
	gerr.Label = "ErrHTTP"
	gerr.Message = reason.Error()
	return
}

func ErrHTTPRespStatus(statusCode int, reason error) (gerr *gerror.GError) {
	gerr = gerror.NewGError()
	gerr.Code = ErrHTTPRespStatusCode
	gerr.Label = "ErrHTTPRespStatus"

	gerr.Message = fmt.Sprintf(`"|Params| %d:%s.\t
								|Reason| %s"`,
		statusCode, http.StatusText(statusCode),
		reason)

	return
}
