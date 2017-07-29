package gonet

import (
	"net"
)

type Reactor struct {
   pipeline *Pipeline
   el	*EventLoop
}



func handle(conn net.Conn){
	r:=&Reactor{NewPipeline(),newEventLoop()}
	b:=make([]byte,1024)
    _,err:=conn.Read(b)
	if err!=nil{
       for {
		   var d interface{}
		   d = b
		   for _, decoder := range r.pipeline.decoders {
			   if d = decoder.onRead(d); d != nil {
				   continue
			   } else {
				   break
			   }
		   }
		   r.el.put(&event{d,})
	   }
	}
}