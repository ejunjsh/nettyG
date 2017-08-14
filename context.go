package gonet

type HandlerContext struct {
	p *Pipeline
    next *HandlerContext
	prev *HandlerContext
	handler Handler
}

func newHandlerContext(p *Pipeline,handler Handler) *HandlerContext {
	return &HandlerContext{p:p,handler:handler}
}

func (h *HandlerContext) FireWrite(data interface{}){
	hc:=h.findNextOutbound()
	if hc!=nil{
		hc.handler.(OutboundHandler).Write(hc,data)
	}
}

func (h *HandlerContext) FireRead(data interface{}){
    hc:=h.findNextInbound()
	if hc!=nil{
		hc.handler.(InboundHandler).Read(hc,data)
	}
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
	next:=h
	for{
		if next.isInbound(){
			return next
		}
		next=next.next
	}
}

func (h *HandlerContext) findNextOutbound() *HandlerContext{
	prev:=h
	for{
		if prev.isOutbound(){
			return prev
		}
		prev=prev.prev
	}
}


