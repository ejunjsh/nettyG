package gonet

type ChannelActiveHandler interface {
	channelActive(c *Context) error
}

type ErrorCaughtHandler interface {
	errorCaught(c *Context,err error)
}

type ChannelReadHandler interface {
	channelRead(c *Context,data interface{}) error
}

type Handler interface {
	ChannelActiveHandler
	ErrorCaughtHandler
	ChannelReadHandler
}

