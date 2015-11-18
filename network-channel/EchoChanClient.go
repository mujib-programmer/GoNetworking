// EchoChanClient
// a client is sent a channel from a server through a shared channel, and uses that private channel.
// This doesn't work directly with network channels: a channel cannot be sent over a network channel. So we have to be a little more
// indirect. Each time a client connects to a server, the server builds new network channels and exports them with new names. Then
// it sends the names of these new channels to the client which imports them. It uses these new channels for communicaiton.
//
// You can execute the EchoChanClient like this:
//		EchoChanClient localhost:2345
//
// Which will output something like this
//		Got importer
//		Imported in
//		Imported out
//		hello  0
//		hello  1
//		hello  2
//		hello  3
//		hello  4
//		hello  5
//		hello  6
//		hello  7
//		hello  8
//		hello  9

package main

import (
	"fmt"
	"golang.org/x/exp/old/netchan"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}

	service := os.Args[1]

	importer, err := netchan.Import("tcp", service)
	checkError(err)

	fmt.Println("Got importer")

	echo := make(chan string)

	importer.Import("echo", echo, netchan.Recv, 1)

	fmt.Println("Imported in")

	count := <-echo

	fmt.Println(count)

	echoIn := make(chan string)

	importer.Import("echoIn"+count, echoIn, netchan.Recv, 1)

	echoOut := make(chan string)

	importer.Import("echoOut"+count, echoOut, netchan.Send, 1)

	for n := 1; n < 10; n++ {

		echoOut <- "hello "

		s := <-echoIn

		fmt.Println(s, n)
	}

	close(echoOut)

	os.Exit(0)
}

func checkError(err error) {

	if err != nil {

		fmt.Println("Fatal error ", err.Error())

		os.Exit(1)
	}
}