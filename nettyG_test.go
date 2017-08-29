package nettyG

import (
	"testing"
	"fmt"
	"os"
	"net"
)


func init(){
	go func() {
		NewBootstrap().Handler(func(channel *channel) {
			channel.Pipeline().AddLast(NewStringCodec()).AddLast(InboundActiveFuc(func(context *HandlerContext) error {
				context.WriteAndFlush("hello netgo")
				return nil
			})).AddLast(InboundReadFuc(func(context *HandlerContext, data interface{}) error {
				if _,ok:=data.(string);ok{
					context.WriteAndFlush("Acknowledge")
					context.Close()
				}
				return nil
			}))
		}).RunServer("tcp",":8981")
	}()
}



func BenchmarkNettyG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", ":8981")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
			os.Exit(1)
		}


		b :=make([]byte,1024)
		conn.Read(b)
		conn.Write([]byte("hello"))
		conn.Read(b)
		conn.Close()
	}
}