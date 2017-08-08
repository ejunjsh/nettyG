package gonet

import "net"

type Client struct {
    pipeline *Pipeline
}

func NewClient() *Client{
	return  &Client{newPipeline()}
}

func (c *Client) AddInboundHandler(handler InboundHandler){
	c.pipeline.addInbound(handler)
}

func (c *Client) AddOutboundHandler(handler OutboundHandler){
	c.pipeline.addOutbound(handler)
}

func (c *Client) Connect(proto string,addr string){
	conn,err:= net.Dial(proto,addr)
	if err!=nil{
		return
	}

	go handle(conn,c.pipeline)
}