package main

import (
	"net"
	"fmt"
	"os"
	//"time"
	"io/ioutil"
)

func main(){
	l,err:=net.Listen("tcp",":9988")

	CheckError(err)


	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn)
	}

}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleConnection(conn net.Conn) {
	//buffer := make([]byte, 1024)

		//n, err := conn.Read(buffer)
		//if err != nil {
		//	Log(conn.RemoteAddr().String(), " connection error: ", err)
		//	return
		//}

		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			// Error Handler
			return
		}

		fmt.Println(string(buf))

		//Log(conn.RemoteAddr().String(), "receive data length:", n)
		//Log(conn.RemoteAddr().String(), "receive data:", buffer[:n])
		//Log(conn.RemoteAddr().String(), "receive data string:", string(buffer[:n]))

}

func Log(v ...interface{}) {
	fmt.Println(v...)
}