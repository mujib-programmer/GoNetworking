// LookupHost
// Utility to perform a DNS lookup on a hostname, and return multiple IP addresses
//
// You can execute the IpMask.go like this:
//    LookupHost www.google.com
//
// Which will output something like this
//    173.194.117.16
//    173.194.117.19
//    173.194.117.18
//    173.194.117.20
//    173.194.117.17
//    2404:6800:4003:805::1010

package main
import (
	"net"
	"os"
	"fmt"
)
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]

	addrs, err := net.LookupHost(name)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}

	os.Exit(0)
}
