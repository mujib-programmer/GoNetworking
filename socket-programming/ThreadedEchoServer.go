// ThreadedEchoServer
// "echo" is another simple IETF service. This just reads what the client data, and sends it back.
// this version of echo server are allowing multiple client to connect using go routine multi threading
//
// You can execute the SimpleEchoServer.go like this:
// 		ThreadedEchoServer
//
// If you run this server, it will just wait there, not doing much. When a client connects to it, it will respond
// by printing the the string data accepted from client to stdout and then return string back as response.
// This server only allow multiple client to connect at the same time.
//
// To testing the server, you can use telnet by this command:
// 		telnet localhost 1201

package main

import (
	"net"
	"os"
	"fmt"
)

func main() {
	service := ":1201" // localhost:1201

	// get *TCPAddr from net type and ip address string given
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// get *TCPListener
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)


	for {
		// get connection from client
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// run as a goroutine to handle multiple client
		go handleClient(conn)
	}
}

// handle client connection
// save all data from client to buffer and then print buffer to stdout
// send back buffer to client as response
func handleClient(conn net.Conn) {
	// close connection on exit
	defer conn.Close()
	var buf [512]byte
	for {
		// read upto 512 bytes
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		// print data to stdout
		fmt.Println(string(buf[0:]))

		// write the n bytes read
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

// checkError print Fatal error to stderr and quit execution if err not nil
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
