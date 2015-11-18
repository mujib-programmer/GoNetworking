// LookupPort
// Utility to get service port number from given service name
//
// You can execute the LookupPort like this:
//    LookupPort tcp telnet
//
// Which will output something like this
//    Service port: 23
//

package main

import (
	"net"
	"os"
	"fmt"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s network-type(tcp or udp) service\n",
			os.Args[0])
		os.Exit(1)
	}

	networkType := os.Args[1]

	service := os.Args[2]

	// lookup port number from given network type and service name
	port, err := net.LookupPort(networkType, service)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	fmt.Println("Service port ", port)


	// The type TCPAddr is a structure containing an IP and a port :
	//
	//     type TCPAddr struct {
    //	       IP IP
	//         Port int
    //     }
	//
	// The function to create a TCPAddr is ResolveTCPAddr
	//    func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
	//
	// where net is one of "tcp", "tcp4" or "tcp6" and the addr is a string composed of a host name or IP address, followed by the port
	// number after a ":", such as "www.google.com:80" or '127.0.0.1:22".

	tcpAddrGoogle, err := net.ResolveTCPAddr("tcp", "www.google.com:80");

	if err != nil {
		fmt.Println("Error getting Google TCP Address: ", err.Error())
		os.Exit(2)
	}

	fmt.Println("Google tcp address are ", tcpAddrGoogle.String())
	fmt.Println("Ip Google are ", tcpAddrGoogle.IP.String())
	fmt.Println("Port Tcp Address are ", tcpAddrGoogle.Port)
	fmt.Println("Google zone tcp address are ", tcpAddrGoogle.Zone)



	os.Exit(0)
}