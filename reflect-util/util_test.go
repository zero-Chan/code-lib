package reflect_util

import (
	"fmt"
	"testing"
)

type HandlerInterface interface {
	Init(id string)
	Handle()
	GetID() string
}

type MyHandler struct {
	Id string
}

func (this *MyHandler) Init(id string) {
	this.Id = id
}

func (this *MyHandler) Handle() {
	fmt.Printf("MyHandler.id=[%d]\n", this.Id)
}

func (this *MyHandler) GetID() string {
	return this.Id
}

func Test_NewInterface(t *testing.T) {
	srchdl := &MyHandler{}
	srchdl.Init("1")

	var srchdlItf HandlerInterface = srchdl
	dsthdlItf, ok := NewInterface(srchdlItf).(HandlerInterface)
	if !ok {
		t.Errorf("convert as src interface type fail. dst type is [%T]", NewInterface(srchdlItf))
		t.FailNow()
	}

	dsthdlItf.Init("2")

	// 不应改变源地址的内容
	if srchdl.GetID() != "1" || srchdl.GetID() == "2" {
		t.Errorf("change src addr.")
		t.FailNow()
	}

	// 能回到 *MyHandler
	dsthdl, ok := dsthdlItf.(*MyHandler)
	if !ok {
		t.Errorf("convert as src type fail. dst type is [%T]", NewInterface(srchdlItf))
		t.FailNow()
	}

	if dsthdl.GetID() != "2" {
		t.Errorf("set id fail.")
		t.FailNow()
	}
}

func Test_CopyInterface(t *testing.T) {
	srchdl := &MyHandler{}
	srchdl.Init("1")

	var srchdlItf HandlerInterface = srchdl
	dsthdlItf, ok := CopyInterface(srchdlItf).(HandlerInterface)
	if !ok {
		t.Errorf("convert as src interface type fail. dst type is [%T]", CopyInterface(srchdlItf))
		t.FailNow()
	}

	if srchdlItf.GetID() != dsthdlItf.GetID() {
		t.Errorf("id not equal. copy fail.")
		t.FailNow()
	}

	dsthdlItf.Init("2")

	// 不应改变源地址的内容
	if srchdl.GetID() != "1" || srchdl.GetID() == "2" {
		t.Errorf("change src addr.")
		t.FailNow()
	}

	// 能回到 *MyHandler
	dsthdl, ok := dsthdlItf.(*MyHandler)
	if !ok {
		t.Errorf("convert as src type fail. dst type is [%T]", CopyInterface(srchdlItf))
		t.FailNow()
	}

	if dsthdl.GetID() != "2" {
		t.Errorf("set id fail.")
		t.FailNow()
	}
}
