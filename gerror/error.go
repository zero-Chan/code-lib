package gerror

import (
	"fmt"
)

type Error interface {
	error
	IsNil() bool
	ErrCode() int64
	ErrorString() string
	ErrorBytes() []byte
}

type GError struct {
	isNil bool

	// 错误码
	Code int64

	// 错误标签
	Label string

	// 错误消息
	Message string

	// 错误等级
	level int64

	// 是否可忽略
	ignore bool
}

func CreateGError() (gerr GError) {
	gerr = GError{
		isNil: false,
	}
	return
}

func NewGError() (gerr *GError) {
	e := CreateGError()
	gerr = &e
	return
}

func (this *GError) Error() (data string) {
	data = fmt.Sprintf("[%d:%s]# %s", this.Code, this.Label, this.Message)
	return
}

func (this *GError) ErrCode() int64 {
	return this.Code
}

func (this *GError) IsNil() bool {
	if this == nil {
		return true
	}

	return this.isNil
}

func (this *GError) ErrorString() (data string) {
	return this.Error()
}

func (this *GError) ErrorBytes() (data []byte) {
	return []byte(this.Error())
}
