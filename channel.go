package gonet

import (
	"net"
	"bytes"
	"time"
	"io"
)

type channel struct {
	conn net.Conn
	pipeline *Pipeline
	writebuffer  *bytes.Buffer
	readbuffer *bytes.Buffer
	flushC chan bool
}

func newChannel(conn net.Conn,pipeline *Pipeline) *channel{
	chl:= &channel{conn,nil,bytes.NewBuffer(make([]byte,0,1024)),bytes.NewBuffer(make([]byte,0,1024)),make(chan bool)}
	chl.pipeline=pipeline
	return chl
}

func (c *channel) Pipeline() *Pipeline{
	return c.pipeline
}

func (c *channel) runReadEventLoop(){
	go func() {
		for{
			b:=make([]byte,1024)
			_,err:=c.conn.Read(b)
			if err!=nil{
				break
			}
			c.pipeline.fireNextRead(b)
		}
	}()
}

func (c *channel) runWriteEventLoop(){
	go func() {
		t:=time.Tick(time.Second)
		for{
			select {
			case <-c.flushC:
				io.Copy(c.conn,c.writebuffer)
			case <-t:
				io.Copy(c.conn,c.writebuffer)
			}
		}
	}()
}

func (c *channel) Write(b []byte){
     c.writebuffer.Write(b)
}

