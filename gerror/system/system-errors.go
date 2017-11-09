package system_err

import (
	"fmt"
	"net/http"

	"code-lib/gerror"
)

// SystemErr 以 10000　为基准
// 增加基数为 100

const (
	ErrSystemCode = 10000

	// HTTPErr
	ErrHTTPCode           = 20100 // HTTP 统一的Error
	ErrHTTPRespStatusCode = 20101 // HTTP 响应Error
	ErrHTTPMuxFilterCode  = 20102 // HTTP 路由过滤Error

	// IOErr
	ErrIOCode      = 20200 // IO　操作统一Error
	ErrIOReadCode  = 20201 // IO Read Error
	ErrIOWriteCode = 20202 // IO Write Error
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

func ErrHTTPMuxFilter(reason error) (gerr *gerror.GError) {
	gerr = gerror.NewGError()
	gerr.Code = ErrHTTPMuxFilterCode
	gerr.Label = "ErrHTTPMuxFilter"
	gerr.Message = reason.Error()
	return
}

// ==========
// IOErr:
// Code: range [20200, 20299)
func ErrIO(reason error) (gerr *gerror.GError) {
	gerr = gerror.NewGError()
	gerr.Code = ErrIOCode
	gerr.Label = "ErrIO"
	gerr.Message = reason.Error()
	return
}

func ErrIORead(src string, reason error) (gerr *gerror.GError) {
	gerr = gerror.NewGError()
	gerr.Code = ErrIOReadCode
	gerr.Label = "ErrIORead"
	gerr.Message = fmt.Sprintf("src[%s], reason: %s", src, reason)
	return
}

func ErrIOWrite(src string, reason error) (gerr *gerror.GError) {
	gerr = gerror.NewGError()
	gerr.Code = ErrIOWriteCode
	gerr.Label = "ErrIOWrite"
	gerr.Message = fmt.Sprintf("src[%s], reason: %s", src, reason)
	return
}
