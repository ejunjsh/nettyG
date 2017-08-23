package netgo

import (
	"testing"
	"fmt"
)

func TestBootstrap_RunServer(t *testing.T) {
	NewBootstrap().Handler(func(channel *channel) {
        channel.Pipeline().AddLast(OutboundWriteFuc(func(context *HandlerContext, data interface{}) error {
			if s,ok:=data.(string);ok{
				context.FireWrite([]byte(s))
				fmt.Println("write 1")
			}
			return nil
		})).AddLast(InboundConnectedFuc(func(context *HandlerContext) error {
			context.FireWrite("hello netgo")
			fmt.Println("channel connected")
			return nil
		})).AddLast(InboundReadFuc(func(context *HandlerContext, data interface{}) error {
			if b,ok:=data.([]byte);ok{
				fmt.Printf("recieve %s",string(b))
			}
			return nil
		}))
	}).RunServer("tcp",":8989")
}