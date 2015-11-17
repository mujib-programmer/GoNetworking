// SimpleEchoServer
// "echo" is another simple IETF service. This just reads what the client types, and sends it back.
// While it works, there is a significant issue with this server: it is single-threaded. While a client has a connection open to it, no other
// client can connect. Other clients are blocked, and will probably time out. Fortunately this is easly fixed by making the client handler
// a go-routine. We have also moved the connection close into the handler, as it now belongs there
//
// You can execute the SimpleEchoServer.go like this:
// 		SimpleEchoServer
//
// If you run this server, it will just wait there, not doing much. When a client connects to it, it will respond
// by printing the string data accepted from client to stdout and then return string back as response.
// This server only allow 1 client to connect at the same time.
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

		handleClient(conn)

		conn.Close() // we're finished
	}
}

// handle client connection
// save all data from client to buffer and then print buffer to stdout
// send back buffer to client as response
func handleClient(conn net.Conn) {

	// buf variable to save data from client
	var buf [512]byte

	for {

		// read length of data from client
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		// print data to stdout
		fmt.Println(string(buf[0:]))

		// send data back to client as response
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