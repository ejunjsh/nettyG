# netgo
a simple netty-like network frame.
````go
	NewBootstrap().Handler(func(channel *channel) {
        channel.Pipeline().AddLast(OutboundWriteFuc(func(context *HandlerContext, data interface{}) error {
			if s,ok:=data.(string);ok{
				context.Write([]byte(s))
			}
			return nil
		})).AddLast(InboundActiveFuc(func(context *HandlerContext) error {
			context.Write("hello netgo")
			fmt.Println("channel connected")
			return nil
		})).AddLast(InboundReadFuc(func(context *HandlerContext, data interface{}) error {
			if b,ok:=data.([]byte);ok{
				fmt.Printf("recieve %s",string(b))
			}
			return nil
		}))
	}).RunServer("tcp",":8989")
````
