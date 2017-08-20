package netgo

import (
	"testing"
	"fmt"
)

func TestBootstrap_RunServer(t *testing.T) {
	NewBootstrap().Handler(func(channel *channel) {
        channel.Pipeline().AddLast(InboundConnectedFuc(func(context *HandlerContext) error {
			context.FireWrite("hello netgo")
			return nil
		})).AddLast(InboundReadFuc(func(context *HandlerContext, data interface{}) error {
			if s,ok:=data.(string);ok{
				fmt.Printf("recieve %s",s)
			}
			return nil
		})).AddLast(OutboundWriteFuc(func(context *HandlerContext, data interface{}) error {
			if s,ok:=data.(string);ok{
				context.FireWrite([]byte(s))
			}
			return nil
		}))
	}).RunServer("tcp",":8989")
}