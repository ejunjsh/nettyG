package gonet



type InboundHandler interface {
	channelActive(c *HandlerContext) error
	channelRead(c *HandlerContext,data interface{}) error
	errorCaught(c *HandlerContext,err error)
}

type Handler interface {
	InboundHandler
}