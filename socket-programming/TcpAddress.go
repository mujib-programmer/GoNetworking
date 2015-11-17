// TcpAddress
// Utility to get tcp address information from given net and addr string
//
// You can execute the IpMask.go like this:
//    TcpAddress tcp www.google.com 80
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

	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s network-type(tcp, tcp4, or tcp6) domain-name(www.google.com) port-number(80, 23, etc) \n",
			os.Args[0])
		os.Exit(1)
	}

	networkType := os.Args[1]

	domainName := os.Args[2]

	portNumber := os.Args[3]

	addrStr := domainName + ":" + portNumber

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

	tcpAddr, err := net.ResolveTCPAddr(networkType, addrStr);

	if err != nil {
		fmt.Println("Error getting ", addrStr, " Address: ", err.Error())
		os.Exit(2)
	}


	fmt.Println(domainName, " tcp address are ", tcpAddr.String())
	fmt.Println("Ip ", domainName, " are ", tcpAddr.IP.String())
	fmt.Println("Port number ", domainName," Tcp Address are ", tcpAddr.Port)
	fmt.Println(domainName, " zone tcp address are ", tcpAddr.Zone)

	os.Exit(0)
}