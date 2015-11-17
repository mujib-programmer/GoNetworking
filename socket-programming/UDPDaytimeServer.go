// UDPDaytimeServer
// The daytime service is very simple and just writes the current time to the client, closes the connection,
// and resumes waiting for the next client.
//
// You can execute the DaytimeServer.go like this:
// 		UDPDaytimeServer
//
// If you run this server, it will just wait there, not doing much. When a client connects to it, it will respond
// by sending the daytime string to it and then return to waiting for the next client.
//
// To testing the server, you can need to use UDPDaytimeClient by this command:
// 		UDPDaytimeClient localhost:1200
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
	service := ":1200" //localhos:1200

	// get *UDPAddr from UDP net type and ip address string given
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	// get *UDPListener
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClient(conn)
	}
}

// handle client connection
// save all data from client to buffer and then print buffer to stdout
// send back buffer to client as response
func handleClient(conn *net.UDPConn) {

	// close connection on exit
	defer conn.Close()

	// buf variable to save data from client
	var buf [512]byte

	// read data from client up to 512 byte buffer
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	fmt.Println("Accept connection from ", conn.LocalAddr())

	daytime := time.Now().String()

	// give response to client
	conn.WriteToUDP([]byte(daytime), addr)
}

// checkError print Fatal error to stderr and quit execution if err not nil
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}