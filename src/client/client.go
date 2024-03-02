package client

import (
	"context"
	"errors"
	"gorpc/src/codec"
	"gorpc/src/common"
	"log"
	"sync"
)

type Client struct {
	cc       codec.Codec
	opt      *common.Option
	header   *codec.Header
	sending  sync.Mutex
	mu       sync.Mutex
	pending  map[uint64]*Call
	seq      uint64
	closing  bool
	shutdown bool
}

func (c *Client) registerCall(call *Call) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closing || c.shutdown {
		return 0, errors.New("rpc client: client closed, register call failed")
	}
	call.Seq = c.seq
	c.pending[call.Seq] = call
	c.seq++
	return call.Seq, nil
}

func (c *Client) removeCall(seq uint64) *Call {
	//TODO implement me
	panic("implement me")
}

func (c *Client) terminateCalls() {
	//TODO implement me
	panic("implement me")
}

func (c *Client) receive() {
	//TODO implement me
	panic("implement me")
}

func (c *Client) send(call *Call) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Call(ctx context.Context, serviceMethod string, args interface{}, reply interface{}) error {
	call := c.Go(serviceMethod, args, reply, make(chan *Call, 1))
	select {
	case <-ctx.Done():
		c.removeCall(call.Seq)
		return errors.New("rpc client: call failed: " + ctx.Err().Error())
	case call := <-call.Done:
		return call.Error
	}
}

func (c *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call, 10)
	} else if cap(done) == 0 {
		log.Panic("rpc client: done channel is unbuffered")
	}
	call := &Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
		Done:          done,
	}
	c.send(call)
	return call
}

func (c *Client) IsAvailable() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return !c.shutdown && !c.closing
}

func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closing {
		return errors.New("rpc client: ErrShutdown")
	}
	return c.cc.Close()
}

var _ clienter = (*Client)(nil)
