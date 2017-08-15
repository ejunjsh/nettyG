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
	c.FireConnected()
	return  nil
}

func (h *headHandler) errorCaught(c *HandlerContext,err error){

}

func (h *headHandler) Write(c *HandlerContext,data interface{}) error{
	b,ok:=data.([]byte)
	if ok{
		c.p.chl.Write(b)
	}
	return nil
}

type tailHandler struct {

}

func (t *tailHandler) errorCaught(c *HandlerContext,err error){

}



func newPipeline() *Pipeline{
     p:=&Pipeline{}
	p.tail=&HandlerContext{p,nil,nil,&tailHandler{}}
	p.head=&HandlerContext{p,nil,nil,&headHandler{}}
	p.head.next=p.tail
	p.tail.prev=p.head
	return p
}



