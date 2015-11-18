// TCPArithClient
// call rpc Arith.Multiply and Arith.Divide on TCPArithServer and print the result.
//
// You can execute the TCPArithClient like this:
//		TCPArithClient localhost:1234
//
// Which will output something like this
//		Arith: 17*8=136
//		Arith: 17/8=2 remainder 1


package main

import (
	"net/rpc"
	"fmt"
	"log"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	
	// TCPArithClient need 1 argument/parameter (host:port)
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	
	service := os.Args[1]
	
	client, err := rpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	
	// Synchronous call
	args := Args{17, 8}
	
	var reply int
	
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
	
	var quot Quotient
	
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}