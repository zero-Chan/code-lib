package notify

type Notify interface {
	Name() string
	Receive()
	Pop() <-chan []byte
	Push([]byte) error
	StopPop()
}
