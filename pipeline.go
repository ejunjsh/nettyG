package netgo

type Pipeline struct {
	head *HandlerContext
	tail *HandlerContext
	chl *channel
}

func (p *Pipeline) fireNextChannelRead(data interface{}){
	p.head.FireChannelRead(data)
}

func (p *Pipeline) fireNextChannelActive(){
	p.head.FireChannelActive()
}



func (p *Pipeline) AddLast(handler Handler) *Pipeline{
    prev:=p.tail.prev
	newH:=newHandlerContext(p,handler)
	newH.prev=prev
	newH.next=p.tail
	prev.next=newH
	p.tail.prev=newH
	return p
}

func (p *Pipeline) AddFirst(handler Handler) *Pipeline{
	next:=p.head.next
	newH:=newHandlerContext(p,handler)
	newH.prev=p.head
	newH.next=next
	p.head.next=newH
	next.prev=newH
	return p
}

type headHandler struct {

}

func (h *headHandler) ChannelRead(c *HandlerContext,data interface{}) error{
	c.FireChannelRead(data)
	return nil
}

func (h *headHandler) ChannelActive(c *HandlerContext) error{
	c.FireChannelActive()
	return  nil
}

func (h *headHandler) ErrorCaught(c *HandlerContext,err error){

}

func (h *headHandler) Write(c *HandlerContext,data interface{}) error{
	b,ok:=data.([]byte)
	if ok{
		c.p.chl.Write(b)
	}
	return nil
}

func (h *headHandler) Close(c *HandlerContext) error{
	return c.p.chl.Close()
}

func (h *headHandler) Flush(c *HandlerContext) error{
	return c.p.chl.Flush()
}

type tailHandler struct {

}

func (t *tailHandler) ErrorCaught(c *HandlerContext,err error){

}



func newPipeline() *Pipeline{
     p:=&Pipeline{}
	p.tail=&HandlerContext{p,nil,nil,&tailHandler{}}
	p.head=&HandlerContext{p,nil,nil,&headHandler{}}
	p.head.next=p.tail
	p.tail.prev=p.head
	return p
}



