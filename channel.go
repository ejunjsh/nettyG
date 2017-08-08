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
}

func newChannel(conn net.Conn,pipeline *Pipeline) *channel{
	return &channel{conn,pipeline,bytes.NewBuffer(make([]byte,0,1024)),bytes.NewBuffer(make([]byte,0,1024))}
}

func (c *channel) read(){
	go func() {
		for{
			b:=make([]byte,1024)
			_,err:=c.conn.Read(b)
			if err!=nil{
				break
			}
			c.pipeline.fireNextWrite(b)
		}
	}()
}

func (c *channel) write(){
	go func() {
		t:=time.Tick(time.Second)
		for{
			select {
			case <-t:
				io.Copy(c.conn,c.writebuffer)
			}
		}
	}()
}

