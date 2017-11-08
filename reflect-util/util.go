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

// 复制一个相同类型的实例指针
// 由于golang语言特性，只能复制public的field, 即首字母大写的field
func CopyInterface(src interface{}) (dst interface{}) {
	srcv := reflect.ValueOf(src)
	srctyp := reflect.Indirect(srcv).Type()

	dstv := reflect.New(srctyp)
	dstvptr := dstv.Elem()

	srcv = srcv.Elem()
	//	fmt.Println("=====type : ", dstvptr.Kind(), dstvptr.Type(), dstvptr.Field(0).Type())

	for i := 0; i < srcv.NumField(); i++ {
		fieldname := srcv.Type().Field(i).Name
		srcfieldval := srcv.Field(i)

		dstfieldval := dstvptr.FieldByName(fieldname)
		if !dstfieldval.IsValid() || !dstfieldval.CanSet() {
			continue
		}

		dstfieldval.Set(srcfieldval)
	}

	return dstv.Interface()
}
