package gonet

type Handler interface {
	onConnected(c *Context) error
	onError(c *Context)
	onClosed(c *Context)
	onMessage(c *Context)
}

