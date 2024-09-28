package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)

	length := len(msg)
	msg = msg[:length-1]
	fmt.Fprintln(conn, "Received '"+msg+"'")
}

func main() {
	ln, _ := net.Listen("tcp", ":8030")
	for {
		conn, _ := ln.Accept()
		go handleConn(conn)
	}
}
