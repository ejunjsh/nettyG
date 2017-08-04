package gonet

import "net"

type HandlerContext struct {
	p *Pipeline
    next *HandlerContext
	handler Handler
}

func newHandlerContext() *HandlerContext {
	return &HandlerContext{}
}

