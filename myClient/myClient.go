package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter a message: ")
		scanner.Scan()
		msg := scanner.Text()
		conn, _ := net.Dial("tcp", "127.0.0.1:8030")
		fmt.Fprintln(conn, msg)

		reader := bufio.NewReader(conn)
		recMsg, _ := reader.ReadString('\n')
		fmt.Println(recMsg)
	}
}
