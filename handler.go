package nettyG



type InboundHandler interface {
	ChannelRead(c *HandlerContext,data interface{}) error
	ChannelActive(c *HandlerContext) error
	Handler
}

type OutboundHandler interface {
	Write(c *HandlerContext,data interface{}) error
	Close(c *HandlerContext) error
	Flush(c *HandlerContext) error
	Handler
}

type Handler interface {
	ErrorCaught(c *HandlerContext,err error)
}


type inbound struct {
	read func (*HandlerContext, interface{}) error
	active func(*HandlerContext) error
}

func (in *inbound) ChannelRead(c *HandlerContext,data interface{}) error{
	if in.read==nil{
		c.FireChannelRead(data)
		return nil
	}
	return in.read(c,data)
}

func (in *inbound)  ChannelActive(c *HandlerContext) error{
	if in.active==nil{
		c.FireChannelActive()
		return nil
	}
	return in.active(c)
}

func (in *inbound) ErrorCaught(c *HandlerContext,err error){

}

func ChannelReadFunc(read func (*HandlerContext, interface{}) error) Handler{
	return &inbound{read:read}
}

func ChannelActiveFunc(active func (*HandlerContext) error) Handler{
	return &inbound{active:active}
}

type outbound struct {
	write func(*HandlerContext, interface{}) error
	flush func(*HandlerContext) error
	close func(*HandlerContext) error
}

func (out *outbound) Write(c *HandlerContext,data interface{}) error{
	if out.write==nil{
		c.Write(data)
		return nil
	}
	return out.write(c,data)
}

func (out *outbound) Flush(c *HandlerContext) error{
	if out.flush==nil{
		c.Flush()
		return nil
	}
	return out.flush(c)
}

func (out *outbound) Close(c *HandlerContext) error{
	if out.close==nil{
		c.Close()
		return nil
	}
	return out.close(c)
}

func (out *outbound) ErrorCaught(c *HandlerContext,err error){

}

func WriteFunc(write func (*HandlerContext,interface{}) error) Handler{
	return &outbound{write:write}
}

func CloseFunc(close func (*HandlerContext) error) Handler{
	return &outbound{close:close}
}

func FlushFunc(flush func (*HandlerContext) error) Handler{
	return &outbound{flush:flush}
}