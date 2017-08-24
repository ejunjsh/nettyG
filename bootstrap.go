package nettyG

import (
	"net"
	"fmt"
)

type ChannelInitializer func(*channel)

type Bootstrap struct{
	initHandler ChannelInitializer
}

func NewBootstrap() *Bootstrap{
	return &Bootstrap{}
}

func (b *Bootstrap) initChannel(conn net.Conn) *channel{
	p:=newPipeline()
	c:=newChannel(conn,p)
	p.chl=c
	b.initHandler(c)
	p.fireNextChannelActive()
	return c
}
func (b *Bootstrap) Handler(handler ChannelInitializer)  *Bootstrap{
	b.initHandler=handler
	return b
}

func (b *Bootstrap) RunServer(proto string,addr string){
	fmt.Printf("%s listen on %s\n",proto,addr)
	l,err:=net.Listen(proto,addr)
	defer l.Close()
	if err!=nil{
		return
	}

	for{
		conn,err:=l.Accept()
		if err!=nil{
			return
		}


		go handle(b.initChannel(conn))
	}
}



//func (b *Bootstrap) RunServer(proto string,addr string){
//	conn,err:= net.Dial(proto,addr)
//	if err!=nil{
//		return
//	}
//
//	go handle(b.initChannel(conn))
//}