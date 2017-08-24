package netgo

import (
	"testing"
	"fmt"
)

func TestBootstrap_RunServer(t *testing.T) {
	NewBootstrap().Handler(func(channel *channel) {
        channel.Pipeline().AddLast(NewStringCodec()).AddLast(InboundActiveFuc(func(context *HandlerContext) error {
			context.Write("hello netgo")
			fmt.Println("channel connected")
			return nil
		})).AddLast(InboundReadFuc(func(context *HandlerContext, data interface{}) error {
			if s,ok:=data.(string);ok{
				fmt.Printf("recieve %s",s)
				context.Write("Acknowledge")
			}
			return nil
		}))
	}).RunServer("tcp",":8989")
}