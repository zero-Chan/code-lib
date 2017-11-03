package controllers

//import (
//	"code-lib/rest"
//)

type BaseController struct {
}

func CreateBaseController() BaseController {
	b := BaseController{}
	return b
}

func NewBaseController() *BaseController {
	b := CreateBaseController()
	return &b
}
