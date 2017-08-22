package netgo


type LineCodec struct {
	delimiter []byte
}

func NewLineCodec(delimiter []byte) *LineCodec {
	return &LineCodec{delimiter}
}

func (l *LineCodec) Read(c *HandlerContext,data interface{}) error{
	return nil
}

func (l *LineCodec) Write(c *HandlerContext,data interface{}) error{
	return nil
}

func (l *LineCodec) Connected(c *HandlerContext) error{
	 c.FireConnected()
	return nil
}
