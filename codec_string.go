package netgo

type StringCodec struct {
	Codec
}


func NewStringCodec() *StringCodec {
	return &StringCodec{}
}

func (l *StringCodec) ChannelRead(c *HandlerContext,data interface{}) error{
	if b,ok:=data.([]byte);ok{
		c.FireChannelRead(string(b))
	}
	return nil
}


func (l *StringCodec) Write(c *HandlerContext,data interface{}) error{
	if s,ok:=data.(string);ok{
		c.Write([]byte(s))
	}
	return nil
}
