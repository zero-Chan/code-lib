package proto

import (
	"fmt"

	"code-lib/gerror"
	encoding_err "code-lib/gerror/encoding"
)

type Protor interface {
	Marshal(v interface{}) ([]byte, gerror.Error)
	Unmarshal(data []byte, v interface{}) gerror.Error
}

type ProtoType int

const (
	ProtoType_JSON = 1
)

var (
	// register here
	protorMap = map[ProtoType]Protor{
		ProtoType_JSON: &JSONProtor{}, // JSON
	}
)

func (this ProtoType) Marshal(v interface{}) (data []byte, gerr gerror.Error) {
	protor, exist := protorMap[this]
	if !exist {
		gerr = encoding_err.ErrEncoding(fmt.Errorf("Unknow ProtoType to Marshal."))
		return
	}

	return protor.Marshal(v)
}

func (this ProtoType) Unmarshal(data []byte, v interface{}) (gerr gerror.Error) {
	protor, exist := protorMap[this]
	if !exist {
		gerr = encoding_err.ErrEncoding(fmt.Errorf("Unknow ProtoType to Unmarshal."))
		return
	}

	return protor.Unmarshal(data, v)
}
