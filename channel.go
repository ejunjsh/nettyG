package nettyG

import (
	"net"
	"bytes"
	//"time"
	//"io"
	"sync"
	//"fmt"
)

type channel struct {
	conn net.Conn
	pipeline *Pipeline
	writebuffer  *bytes.Buffer
	readbuffer *bytes.Buffer
	flushC chan bool
	closeC chan struct{}
	writeLocker sync.Mutex
}

func newChannel(conn net.Conn,pipeline *Pipeline) *channel{
	chl:= &channel{conn,nil,bytes.NewBuffer(make([]byte,0,1024)),
		bytes.NewBuffer(make([]byte,0,1024)),make(chan bool,1),make(chan struct{}),sync.Mutex{}}
	chl.pipeline=pipeline
	return chl
}

func (c *channel) Pipeline() *Pipeline{
	return c.pipeline
}

func (c *channel) runReadEventLoop(){
	//go func() {
		for{
			b:=make([]byte,1024)
			n,err:=c.conn.Read(b)
			if err!=nil{
				//c.closeC<- struct{}{}
				return
			}
			c.pipeline.fireNextChannelRead(b[:n])
		}
	//}()
}

//func (c *channel) runWriteEventLoop(){
//	go func() {
//		t:=time.Tick(time.Millisecond)
//		for{
//			select {
//			case <-c.flushC:
//				c.writeLocker.Lock()
//				io.Copy(c.conn,c.writebuffer)
//				c.writeLocker.Unlock()
//			case <-c.closeC:
//				return
//			case <-t:
//				c.writeLocker.Lock()
//				_,err:=io.Copy(c.conn,c.writebuffer)
//				if err!=nil{
//					fmt.Println(err)
//					return
//				}
//				t=time.Tick(time.Millisecond)
//				c.writeLocker.Unlock()
//			}
//		}
//	}()
//}

//func (c *channel) Write(b []byte){
//	c.writeLocker.Lock()
//	defer c.writeLocker.Unlock()
//	c.writebuffer.Write(b)
//}

func (c *channel) Write(b []byte){
	c.conn.Write(b)
}

func (c *channel) Close() error{
	return c.conn.Close()
}

func (c *channel) Flush() error{
	//c.flushC<-true
	return nil
}