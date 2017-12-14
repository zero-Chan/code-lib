package rest

import (
	"github.com/pborman/uuid"
)

type Cmd struct {
	ID string
}

func CreateCmd() Cmd {
	ch := Cmd{
		ID: uuid.New(),
	}

	return ch
}

func NewCmd() *Cmd {
	ch := CreateCmd()
	return &ch
}
