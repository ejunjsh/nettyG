package netgo



type InboundHandler interface {
	Read(c *HandlerContext,data interface{}) error
	Connected(c *HandlerContext) error
	Handler
}

type OutboundHandler interface {
	Write(c *HandlerContext,data interface{}) error
	Handler
}

type Handler interface {
	ErrorCaught(c *HandlerContext,err error)
}


type inbound struct {
	read func (*HandlerContext, interface{}) error
	connected func(*HandlerContext) error
}

func (in *inbound) Read(c *HandlerContext,data interface{}) error{
	if in.read==nil{
		c.FireRead(data)
		return nil
	}
	return in.read(c,data)
}

func (in *inbound)  Connected(c *HandlerContext) error{
	if in.connected==nil{
		c.FireConnected()
		return nil
	}
	return in.connected(c)
}

func (in *inbound) ErrorCaught(c *HandlerContext,err error){

}

func InboundReadFuc(read func (*HandlerContext, interface{}) error) Handler{
	return &inbound{read:read}
}

func InboundConnectedFuc(connected func (*HandlerContext) error) Handler{
	return &inbound{connected:connected}
}

type outbound struct {
	write func(*HandlerContext, interface{}) error
}

func (out *outbound) Write(c *HandlerContext,data interface{}) error{
	if out.write==nil{
		c.FireWrite(data)
		return nil
	}
	return out.write(c,data)
}

func (out *outbound) ErrorCaught(c *HandlerContext,err error){

}

func OutboundWriteFuc(write func (*HandlerContext,interface{}) error) Handler{
	return &outbound{write:write}
}
