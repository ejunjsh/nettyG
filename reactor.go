package gonet

import (
	"net"
)


func handle(conn net.Conn,p *Pipeline){
	chl:=newChannel(conn,p)
	chl.runReadEventLoop()
	chl.runWriteEventLoop()
}