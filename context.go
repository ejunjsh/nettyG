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

func (h *HandlerContext) Write(data interface{}){

}

func (h *HandlerContext) FireRead(data interface{}){

}

func (h *HandlerContext) isInbound() bool{
	_,ok:= h.handler.(InboundHandler)
	return  ok
}

func (h *HandlerContext) isOutbound() bool{
	_,ok:= h.handler.(OutboundHandler)
	return ok
}

func (h *HandlerContext) findNextInbound() *HandlerContext{
	next:=h.next
	for{
		if h.next==nil{
			return nil
		}
		if next.isInbound(){
			return next
		}
	}
}

func (h *HandlerContext) findNextOutbound() *HandlerContext{
	next:=h.next
	for{
		if h.next==nil{
			return nil
		}
		if next.isOutbound(){
			return next
		}
	}
}


