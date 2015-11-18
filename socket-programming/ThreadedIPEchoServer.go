// ThreadedIPEchoServer
// "echo" is another simple IETF service. This just reads what the client data, and sends it back.
// this version of echo server are allowing multiple client to connect using go routine multi threading
// ThreadedIPEchoServer accept tcp or udp connection
//
// You can execute the ThreadedIPEchoServer like this:
// 		ThreadedIPEchoServer
//
// If you run this server, it will just wait there, not doing much. When a client connects to it, it will respond
// by printing the the string data accepted from client to stdout and then return string back as response.
// This server only allow multiple client to connect at the same time.
//
// To testing the server, you can use telnet by this command:
// 		telnet localhost 1200

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	service := ":1200" // localhost:1200

	// net.Listen accep "tcp", "tcp4", "tcp6", "udp", "udp4" or "udp6"
	listener, err := net.Listen("tcp", service)
	checkError(err)

	for {

		// get conn from client
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// using goroutine to handle client connection. every client connection will handled at separate thread
		go handleClient(conn)
	}
}

// handle client connection
// save all data from client to buffer and then print buffer to stdout
// send back buffer to client as response
func handleClient(conn net.Conn) {

	// close connection at exit
	defer conn.Close()

	// buf var to read data from client
	var buf [512]byte

	for {
		// read data from client
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		// write data from client as response
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}

	}
}

// return result byte buffer
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
