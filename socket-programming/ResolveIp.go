// ResolveIP
// Utility to get information about Ip Address from given domain name string
//
// You can execute the ResolveIP like this:
//    ResolveIP www.google.com
//
// Which will output something like this
//    Resolved address is  114.125.1.57

package main

import (
	"net"
	"os"
	"fmt"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}

	name := os.Args[1]

	// perform dns lookup on IP host name
	// The function ResolveIPAddr will perform a DNS lookup on a hostname, and return a single IP address
	addr, err := net.ResolveIPAddr("ip", name)

	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}

	fmt.Println("Resolved address is ", addr.String())

	os.Exit(0)
}