package gonet

import "net"

type ChannelContext struct {
    m map[string]interface{}
	p *Pipeline
	conn net.Conn
}

func newContext() *ChannelContext {
	return &ChannelContext{}
}

func (c *ChannelContext) Add(key string,data interface{}){
    c.m[key]=data
}


func (c *ChannelContext) Write(data interface{}){
	d:=data
	for _, encoder := range c.p.encoders {
		if d = encoder.onWrite(data); d != nil {
			continue
		} else {
			break
		}
	}

	if b,ok:=d.([]byte);ok{
		c.conn.Write(b)
	}else {

	}
}
