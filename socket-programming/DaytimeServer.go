// DaytimeServer
// The daytime service is very simple and just writes the current time to the client, closes the connection,
// and resumes waiting for the next client.
//
// You can execute the DaytimeServer.go like this:
// 		DaytimeServer
//
// If you run this server, it will just wait there, not doing much. When a client connects to it, it will respond
// by sending the daytime string to it and then return to waiting for the next client.
//
// To testing the server, you can use telnet by this command:
// 		telnet localhost 1200
//
// Server will print: Accept connection from  127.0.0.1:1200

package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200" // localhost:1200

	// get *TCPAddr from net type and ip address string given
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// get *TCPListener,
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		// get connection from client
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("Accept connection from ", conn.LocalAddr())

		daytime := time.Now().String()

		// give response to client
		conn.Write([]byte(daytime)) // don't care about return value

		// close connection
		conn.Close()

		// we're finished with this client
	}
}

// checkError print Fatal error to stderr and quit execution if err not nil
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}