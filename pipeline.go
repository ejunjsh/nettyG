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


func (p *Pipeline) addInbound(inbound InboundHandler){
	next:=p.head
   for{
	  if next.next==nil{
          next=newHandlerContext(p,inbound)
		  return
	  } else {
	      next=next.next
	  }
   }
}

func (p *Pipeline) addOutbound(outbound OutboundHandler){
	next:=p.tail
	for{
		if next.next==nil{
			next=newHandlerContext(p,outbound)
			return
		} else {
			next=next.next
		}
	}
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
	p.tail=&HandlerContext{p,nil,&tailHandler{}}
	p.head=&HandlerContext{p,nil,&headHandler{}}
	return p
}

func newPipelineWithChannel(pl *Pipeline,chl *channel) *Pipeline{
	p:=&Pipeline{}
	p.tail=pl.tail
	p.head=pl.head
	p.chl=chl
	return p
}


