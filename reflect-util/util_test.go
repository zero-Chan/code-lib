package reflect_util

import (
	"fmt"
	"testing"
)

type HandlerInterface interface {
	Init(id string)
	Handle()
	ID() string
}

type MyHandler struct {
	id string
}

func (this *MyHandler) Init(id string) {
	this.id = id
}

func (this *MyHandler) Handle() {
	fmt.Printf("MyHandler.id=[%d]\n", this.id)
}

func (this *MyHandler) ID() string {
	return this.id
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
	if srchdl.ID() != "1" || srchdl.ID() == "2" {
		t.Errorf("change src addr.")
		t.FailNow()
	}

	// 能回到 *MyHandler
	dsthdl, ok := dsthdlItf.(*MyHandler)
	if !ok {
		t.Errorf("convert as src type fail. dst type is [%T]", NewInterface(srchdlItf))
		t.FailNow()
	}

	if dsthdl.ID() != "2" {
		t.Errorf("set id fail.")
		t.FailNow()
	}
}
