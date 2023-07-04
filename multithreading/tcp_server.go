package multithreading

import (
	"fmt"
	"net"
	"time"
)

func do(conn net.Conn) {
	request := make([]byte, 1024)
	conn.Read(request)
	time.Sleep(5 * time.Second)
	fmt.Println("Processing the request")
	str := []byte("HTTP/1.1 200 OK\r\n\r\nHey how you doiing, World!\r\n")
	conn.Write(str)
	conn.Close()
}

func Server() {
	port := ":1234"
	listener, err := net.Listen("tcp", port) // Listening on port
	defer listener.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Listening on port 1234")
	for {
		conn, err := listener.Accept() // Accepting connections , blocking call
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Connection accepted")
		go do(conn) // Handling connection in go routines
	}
}
