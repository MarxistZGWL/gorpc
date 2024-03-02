package client

import "context"

type clienter interface {
	registerCall(call *Call) (uint64, error)
	removeCall(seq uint64) *Call
	terminateCalls(err error)
	receive()
	send(call *Call)
	Call(ctx context.Context, serviceMethod string, args interface{}, reply interface{}) error
	Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call
	IsAvailable() bool
	Close() error
}
