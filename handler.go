package gonet

type ChannelActiveHandler interface {
	channelActive(c *ChannelContext) error
}

type ErrorCaughtHandler interface {
	errorCaught(c *ChannelContext,err error)
}

type ChannelReadHandler interface {
	channelRead(c *ChannelContext,data interface{}) error
}

type Handler interface {
	ChannelActiveHandler
	ErrorCaughtHandler
	ChannelReadHandler
}

