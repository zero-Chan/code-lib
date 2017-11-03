package controllers

import (
	"code-lib/rest"

	"code-lib/rest/http/example/server-test/controllers"
)

type ClientAdd struct {
	*controllers.BaseController
}

func CreateClientAdd() ClientAdd {
	c := ClientAdd{
		BaseController: controllers.NewBaseController(),
	}

	return c
}

func NewClientAdd() *ClientAdd {
	c := CreateClientAdd()
	return &c
}

func Preprocess() (resp *rest.RestResponse) {
	resp = rest.NewRestResponse()

	return
}

func Handler() (resp *rest.RestResponse) {
	resp = rest.NewRestResponse()

	return
}
