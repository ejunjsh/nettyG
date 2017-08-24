package netgo


type LineCodec struct {
	delimiter []byte
}

func NewLineCodec(delimiter []byte) *LineCodec {
	return &LineCodec{delimiter}
}

func (l *LineCodec) ChannelRead(c *HandlerContext,data interface{}) error{
	return nil
}

func (l *LineCodec) Write(c *HandlerContext,data interface{}) error{
	return nil
}

func (l *LineCodec) ChannelActive(c *HandlerContext) error{
	 c.FireChannelActive()
	return nil
}
