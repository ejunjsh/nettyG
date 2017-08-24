package nettyG

type Codec struct {

}

func (c *Codec) ChannelActive(h *HandlerContext) error{
	h.FireChannelActive()
	return nil
}

func (c *Codec) Close(h *HandlerContext) error{
	h.Close()
	return nil
}

func (c *Codec) Flush(h *HandlerContext) error{
	h.Flush()
	return nil
}

func (c *Codec) ErrorCaught(h *HandlerContext,err error){

}