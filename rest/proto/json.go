package proto

import (
	"encoding/json"

	"code-lib/gerror"
	encoding_err "code-lib/gerror/encoding"
)

type JSONProtor struct {
}

func (this *JSONProtor) Marshal(v interface{}) (data []byte, gerr gerror.Error) {
	var (
		err error
	)

	data, err = json.Marshal(v)
	if err != nil {
		gerr = encoding_err.ErrJSONMarshal(err)
		return
	}

	return
}

func (this *JSONProtor) Unmarshal(data []byte, v interface{}) (gerr gerror.Error) {
	var (
		err error
	)

	err = json.Unmarshal(data, v)
	if err != nil {
		gerr = encoding_err.ErrJSONUnmarshal(err)
		return
	}

	return
}
