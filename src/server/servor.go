package server

import (
	codec2 "gorpc/src/codec"
	"gorpc/src/common"
	"io"
	"net"
	"sync"
	"time"
)

type servor interface {
	Register(rcvr interface{}) error
	findService(serviceMethod string) (svc *service, mtype *methodType, err error)
	Accept(lis net.Listener)
	serveConn(conn io.ReadWriteCloser)
	serveCodec(cc codec2.Codec, opt *common.Option)
	readRequest(cc codec2.Codec) (*request, error)
	readRequestHeader(cc codec2.Codec) (*codec2.Header, error)
	handleRequest(cc codec2.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup, duration time.Duration)
	sendResponse(cc codec2.Codec, header *codec2.Header, body interface{}, sending *sync.Mutex)
}
