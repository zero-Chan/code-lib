package encoding_err

import (
	"code-lib/gerror"
)

// EncodingErr 以 20000　为基准

const (
	ErrEncodingCode = 20000

	// JSONErr
	ErrJSONCode          = 20100
	ErrJSONMarshalCode   = 20101
	ErrJSONUnmarshalCode = 20102
)

// =====
// EncodingErr:
// Code: range [20000, 30000)
func ErrEncoding(reason error) (gerr *gerror.GError) {
	gerr = gerror.NewGError()
	gerr.Code = ErrEncodingCode
	gerr.Label = "ErrEncoding"
	gerr.Message = reason.Error()
	return
}

// ==========
// JSONErr:
// Code: range [20100, 20199)
func ErrJSON(reason error) (gerr *gerror.GError) {
	gerr = gerror.NewGError()
	gerr.Code = ErrJSONCode
	gerr.Label = "ErrJSON"
	gerr.Message = reason.Error()
	return
}

func ErrJSONMarshal(reason error) (gerr *gerror.GError) {
	gerr = gerror.NewGError()
	gerr.Code = ErrJSONMarshalCode
	gerr.Label = "ErrJsonMarshal"
	gerr.Message = reason.Error()
	return
}

func ErrJSONUnmarshal(reason error) (gerr *gerror.GError) {
	gerr = gerror.NewGError()
	gerr.Code = ErrJSONUnmarshalCode
	gerr.Label = "ErrJSONUnmarshal"
	gerr.Message = reason.Error()
	return
}

// ==========
