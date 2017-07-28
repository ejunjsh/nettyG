package gonet

import (
	"bytes"
	"net"
	"io"
)

type BufferChannel struct{
	 buffer *bytes.Buffer
	conn net.Conn
}

func newBufferChannel(conn net.Conn) *BufferChannel{
	return &BufferChannel{bytes.NewBuffer(make([]byte,0)),conn}
}


func (b *BufferChannel) Read(){
    b.buffer.Read()
}

func (b *BufferChannel) Write(){
	b.buffer.Write()
}

func (b *BufferChannel) Flush(){
	io.Copy(b.conn,b.buffer)
}