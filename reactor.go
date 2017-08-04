package gonet

import (
	"net"
	"bytes"
)


func handle(conn net.Conn,p *Pipeline){
	for  {
		writeBuffer:= bytes.NewBuffer(make([]byte,0,1024*10))
		read:=func(){
			b:=make([]byte,1024)
			n,error:=conn.Read(b)
			p.head.handler.channelRead(p.head,b)
		}
		write:=func(data interface{}){
			if ok,b:= data.(bytes.Buffer);ok{
			}
		}

		read()
		write()
	}
}