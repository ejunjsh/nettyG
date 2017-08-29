package nettyG

import (
	"net"
	"testing"
	"fmt"
	"os"
)

func init() {
	go func() {
		l,err:=net.Listen("tcp",":9988")

		if err!=nil{
			return
		}

		defer l.Close()

		for {
			conn, err := l.Accept()
			if err != nil {
				continue
			}
			go func() {
				conn.Write([]byte("hello gostd"))
				b :=make([]byte,1024)
				conn.Read(b)
				conn.Write([]byte("Acknowledge"))
				conn.Close()
			}()
		}
	}()
}



func BenchmarkGostd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", ":9988")
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