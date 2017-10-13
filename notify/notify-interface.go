package notify

type Notify interface {
	Init() error
	Name() string
	Receive() error
	Pop() <-chan []byte
	Push([]byte) error
	StopPop()
	Ack() error
}
