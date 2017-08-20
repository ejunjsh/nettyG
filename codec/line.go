package codec

import  "github.com/ejunjsh/netgo"

type LineCodec struct {
	delimiter string
}

func NewLineCodec(delimiter string) *LineCodec {
	return &LineCodec{delimiter}
}

func (l *LineCodec) Read(c *netgo.HandlerContext,data interface{}) error{
	return nil
}

func (l *LineCodec) Write(c *netgo.HandlerContext,data interface{}) error{
	return nil
}

func (l *LineCodec) Connected(c *netgo.HandlerContext) error{
	 c.FireConnected()
	return nil
}
