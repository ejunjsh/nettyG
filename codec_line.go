package netgo


type LineCodec struct {
	delimiter []byte
	Codec
}

func NewLineCodec(delimiter []byte) *LineCodec {
	return &LineCodec{delimiter:delimiter}
}

func (l *LineCodec) ChannelRead(c *HandlerContext,data interface{}) error{
	return nil
}

func (l *LineCodec) Write(c *HandlerContext,data interface{}) error{
	return nil
}

