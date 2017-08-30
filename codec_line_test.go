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
				AddLast(NewLineCodec("\n")).
				AddLast(NewStringCodec()).
				AddLast(InboundActiveFuc(func(context *HandlerContext) error {
				context.WriteAndFlush("hello netgo")
				return nil
			})).AddLast(InboundReadFuc(func(context *HandlerContext, data interface{}) error {
				if d,ok:=data.(string);ok{
					context.Write(d)
				}
				return nil
			}))
		}).RunServer("tcp",":8981")
	}()

	conn, err := net.Dial("tcp", ":8981")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}


	b :=make([]byte,1024)
	n,_:=conn.Read(b)
	conn.Write([]byte("hellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohello\n world\n"))
	n,_=conn.Read(b)
	fmt.Print(string(b[0:n]))
	n,_=conn.Read(b)
	fmt.Print(string(b[0:n]))
	conn.Close()
	time.Sleep(1000*time.Second)
}