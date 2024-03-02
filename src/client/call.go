package client

type Call struct {
	ServiceMethod string
	Args          interface{}
	Reply         interface{}
	Error         error
	Seq           uint64
	Done          chan *Call
}

func (call *Call) done() {
	call.Done <- call
}
