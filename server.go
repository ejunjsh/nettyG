package gonet

import "net"

type Server struct {
    pipeline *Pipeline
}

func NewServer() *Server{
	return &Server{newPipeline()}
}

func (s *Server) AddInboundHandler(handler InboundHandler){
	s.pipeline.addInbound(handler)
}

func (s *Server) AddOutboundHandler(handler OutboundHandler){
	s.pipeline.addOutbound(handler)
}

func (s *Server) Run(proto string,addr string){
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

		go handle(conn,s.pipeline)
	}
}
