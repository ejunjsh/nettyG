package gonet



type InboundHandler interface {
	Read(c *HandlerContext,data interface{}) error
	Connected(c *HandlerContext) error
	Handler
}

type OutboundHandler interface {
	Write(c *HandlerContext,data interface{}) error
	Handler
}

type Handler interface {
	errorCaught(c *HandlerContext,err error)
}