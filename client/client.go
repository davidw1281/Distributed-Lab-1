package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

func read(conn net.Conn, wg sync.WaitGroup) {
	//TODO In a continuous loop, read a message from the server and display it.
	for {
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
	}
	defer wg.Done()
}

func write(conn net.Conn, wg sync.WaitGroup) {
	//TODO Continually get input from the user and send messages to the server.
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter a message: ")
		scanner.Scan()
		msg := scanner.Text()
		fmt.Fprintln(conn, msg)
	}
	defer wg.Done()
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()

	var wg sync.WaitGroup
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
	conn, _ := net.Dial("tcp", *addrPtr)
	wg.Add(2)
	go write(conn, wg)
	go read(conn, wg)
	wg.Wait()
}
