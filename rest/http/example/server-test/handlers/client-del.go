package handlers

type ClientDel struct {
}

func CreateClientDel() ClientDel {
	cli := ClientDel{}

	return cli
}

func NewClientDel() *ClientDel {
	cli := CreateClientDel()
	return &cli
}
