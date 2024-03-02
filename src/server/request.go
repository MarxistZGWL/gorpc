package server

import (
	"gorpc/src/codec"
	"reflect"
)

type request struct {
	h            *codec.Header
	argv, replyv reflect.Value
	svc          *service
	mtype        *methodType
}
