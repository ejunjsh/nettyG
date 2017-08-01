package gonet

import "net"

type Server struct {

}

func NewServer() *Server{
	return &Server{}
}

func (s *Server) Run(proto string,addr string){
	l,err:=net.Listen(proto,addr)
	defer l.Close()
	if err!=nil{
		return
	}

	r:=&Reactor{pipeline:NewPipeline()}
	r.el.run(r.pipeline)

	for{
		conn,err:=l.Accept()
		if err!=nil{
			return
		}

		go handle(conn,r)
	}
}
