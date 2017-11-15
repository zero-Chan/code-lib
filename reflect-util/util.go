package reflect_util

import (
	//	"fmt"
	"reflect"
)

// 创建一个相同类型的实例的指针
// e.g: src is int64, dst is *int64
// e.g: src is WriterInterface from FileStruct, dst is *FileStruct
func NewInterface(src interface{}) (dst interface{}) {
	srcv := reflect.ValueOf(src)
	srctyp := reflect.Indirect(srcv).Type()

	dstv := reflect.New(srctyp)
	return dstv.Interface()
}
