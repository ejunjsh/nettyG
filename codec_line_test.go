package nettyG

import (
	"testing"
	"net"
	"fmt"
	"os"
	"time"
)

func TestNewLineCodec(t *testing.T) {
	go func() {
		NewBootstrap().Handler(func(channel *Channel) {
			channel.Pipeline().
				AddLast(NewLineCodec("\r\n\r\n")).
				AddLast(NewStringCodec()).
				AddLast(ChannelActiveFunc(func(context *HandlerContext) error {
				context.WriteAndFlush("hello nettyG")
				return nil
			})).AddLast(ChannelReadFunc(func(context *HandlerContext, data interface{}) error {
				if d,ok:=data.(string);ok{
					context.Write(d)
				}
				return nil
			}))
		}).RunServer("tcp",":8981")
	}()
    time.Sleep(2*time.Second)
	conn, err := net.Dial("tcp", ":8981")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}


	b :=make([]byte,1024)
	n,_:=conn.Read(b)
	conn.Write([]byte("hellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohello\r\n\r\nworld\r\n\r\n"))
	n,_=conn.Read(b)
	fmt.Print(string(b[0:n]))
	n,_=conn.Read(b)
	fmt.Print(string(b[0:n]))
	conn.Close()
	time.Sleep(1000*time.Second)
}
