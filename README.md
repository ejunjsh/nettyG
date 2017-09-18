# nettyG
[![Build Status](https://travis-ci.org/ejunjsh/nettyG.svg?branch=master)](https://travis-ci.org/ejunjsh/nettyG)

a simple netty-like network framework.
````go
NewBootstrap().Handler(func(channel *Channel) {
			channel.Pipeline().
				AddLast(NewLineCodec("\r\n\r\n")).
				AddLast(NewStringCodec()).
				AddLast(ChannelActiveFunc(func(context *HandlerContext) error {
				context.WriteAndFlush("hello netgo")
				return nil
			})).AddLast(ChannelReadFunc(func(context *HandlerContext, data interface{}) error {
				if d,ok:=data.(string);ok{
					context.Write(d)
				}
				return nil
			}))
		}).RunServer("tcp",":8981")
````

benchmark
nettyG vs go standard lib
````bash
$ cd bench
$ go test -bench .
tcp listen on :8981
BenchmarkGostd-8           10000            123746 ns/op
BenchmarkNettyG-8          10000            134753 ns/op
PASS
ok      github.com/ejunjsh/nettyG       2.623s

````
