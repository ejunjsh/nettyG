package gonet

type Handler interface {
	channelActive(c *Context) error
	exceptionCaught(c *Context,err error)
	channelRead(c *Context,data interface{}) error
}

