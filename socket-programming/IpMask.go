// Ip Mask
// Utility to get information about Ip Mask from given Ip Address string
//
// You can execute the IpMask.go like this:
//    IpMask 127.0.0.1
//
// Which will output something like this
//    Address is  127.0.0.1
//    Default mask length is  32
//    Leading ones count is  8
//    Mask is (hex)  ff000000
//    Network is  127.0.0.0
//

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	dotAddr := os.Args[1]

	addr := net.ParseIP(dotAddr)

	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	// mask is the default mask returned by a method of IP.DefaultMask()
	mask := addr.DefaultMask()

	// A mask can then be used by a method of an IP address to find the network for that IP address
	network := addr.Mask(mask)

	ones, bits := mask.Size()

	fmt.Println("Address is ", addr.String())
	fmt.Println("Default mask length is ", bits)
	fmt.Println("Leading ones count is ", ones)
	fmt.Println("Mask is (hex) ", mask.String())
	fmt.Println("Network is ", network.String())
	os.Exit(0)
}