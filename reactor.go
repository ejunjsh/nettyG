package gonet

import (
	"net"
	"bytes"
)

type Reactor struct {

}

func New() *Reactor{

}

func Run(proto string,addr string){
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

		go handle(conn)
	}
}

func handle(conn net.Conn){
	buffer:=bytes.NewBuffer(make([]byte,100))

}