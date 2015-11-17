// UDPDaytimeClient
// The daytime client connect to UDPDaytimeServer
//
// You can execute the UDPDaytimeClient.go like this:
// 		UDPDaytimeClient localhost:1200
//
// Which will output something like this
//      2015-11-17 06:30:48.690420607 +0700 WIB

package main

import (
	"net"
	"os"
	"fmt"
)

func main() {

	// UDPDaytimeClient need 1 argument/parameter to execute
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	// accept first parameter (ip-address:port string) as service
	service := os.Args[1]

	// get *UDPAddr from UDP net type and ip address string given
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	// get *UDPConn
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	// send data to server. Data will be ignored by server, so we can send anything data
	_, err = conn.Write([]byte("anything"))
	checkError(err)

	// create 512 byte buffer var to save response data from server
	var buf [512]byte

	// read data response from server
	n, err := conn.Read(buf[0:])
	checkError(err)

	// print response to stdout
	fmt.Println(string(buf[0:n]))

	// close application
	os.Exit(0)
}

// checkError print Fatal error to stderr and quit execution if err not nil
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}