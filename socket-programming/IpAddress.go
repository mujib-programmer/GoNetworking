// Ip Address
// Utility to get information about Ip Address from given Ip Address string
//
// You can execute the IpAddress like this:
//    IpAddress 127.0.0.1
//
// Which will output something like this
//    The address is  127.0.0.1
//
package main

import (
	"net"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	// the function ParseIP(String) will take a dotted IPv4 address or a colon IPv6 address
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		// the IP method String will return a string.
		// Note that you may not get back what you started with: the string form of 0:0:0:0:0:0:0:1 is ::1.
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}