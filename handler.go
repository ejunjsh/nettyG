package gonet



type InboundHandler interface {
	Read(c *HandlerContext) error
	Connected(c *HandlerContext,data interface{}) error
	Handler
}

type OutboundHandler interface {
	Write(c *HandlerContext) error
	Handler
}

type Handler interface {
	errorCaught(c *HandlerContext,err error)
}