package rest

import (
	"github.com/pborman/uuid"
)

type RestChannel struct {
	ID string
}

func CreateRestChannel() RestChannel {
	ch := RestChannel{
		ID: uuid.New(),
	}

	return ch
}

func NewRestChannel() *RestChannel {
	ch := CreateRestChannel()
	return &ch
}
