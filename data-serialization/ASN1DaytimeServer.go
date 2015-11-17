// ASN1DaytimeServer
// The daytime service is very simple and just writes the current time to the client, closes the connection,
// and resumes waiting for the next client.
//
// You can execute the DaytimeServer.go like this:
// 		ASN1DaytimeServer
//
// If you run this server, it will just wait there, not doing much. When a client connects to it, it will respond
// by sending the daytime string to it and then return to waiting for the next client.
//
// To testing the server, you can use telnet by this command:
// 		ASN1DaytimeClient localhost:1200
//
// Server will print: Accept connection from  127.0.0.1:1200
// This client and server are exchanging ASN.1 encoded data values, not textual strings.

package main

import (
	"encoding/asn1"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200" // localhost:1200

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("Accept connection from ", conn.LocalAddr())

		daytime := time.Now()

		// Ignore return network errors.
		mdata, _ := asn1.Marshal(daytime)
		conn.Write(mdata)
		conn.Close() // we're finished
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}