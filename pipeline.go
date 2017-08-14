package gonet


type Pipeline struct {
	head *HandlerContext
	tail *HandlerContext
	chl *channel
}

func (p *Pipeline) fireNextRead(data interface{}){
	p.head.FireRead(data)
}

func (p *Pipeline) fireNextWrite(data interface{}){
	p.tail.FireWrite(data)
}


func (p *Pipeline) AddLast(handler Handler){
    prev:=p.tail.prev
	newH:=newHandlerContext(p,handler)
	newH.prev=prev
	newH.next=p.tail
	prev.next=newH
	p.tail.prev=newH
}

func (p *Pipeline) AddFirst(handler Handler){
	next:=p.head.next
	newH:=newHandlerContext(p,handler)
	newH.prev=p.head
	newH.next=next
	p.head.next=newH
	next.prev=newH
}

type headHandler struct {

}

func (h *headHandler) Read(c *HandlerContext,data interface{}) error{
	c.FireRead(data)
	return nil
}

func (h *headHandler) Connected(c *HandlerContext) error{

    return nil
}

func (h *headHandler) errorCaught(c *HandlerContext,err error){

}

type tailHandler struct {

}

func (t *tailHandler) errorCaught(c *HandlerContext,err error){

}

func (t *tailHandler) Write(c *HandlerContext,data interface{}) error{
	c.FireWrite(data)
	return nil
}

func newPipeline() *Pipeline{
     p:=&Pipeline{}
	p.tail=&HandlerContext{p,nil,nil,&tailHandler{}}
	p.head=&HandlerContext{p,nil,nil,&headHandler{}}
	p.head.next=p.tail
	p.tail.prev=p.head
	return p
}

func newPipelineWithChannel(pl *Pipeline,chl *channel) *Pipeline{
	p:=&Pipeline{}
	p.tail=pl.tail
	p.head=pl.head
	p.chl=chl
	return p
}


