// JSONEchoServer
// decode json marshalled data from client, and then encode back data to sent as response to client
//
// You can execute the ASN1DaytimeClient like this:
// 		JSONEchoServer
//
// Which will output something like this
//      Client Newmarch
//		home: jan@newmarch.name
//		work: j.newmarch@boxhill.edu.au



package main

import (
	"fmt"
	"net"
	"os"
	"encoding/json"
)

type Person struct {
	Name Name
	Email []Email
}

type Name struct {
	Family string
	Personal string
}

type Email struct {
	Kind string
	Address string
}

// String() method for person struct
func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

func main() {

	// create person data
	personServer := Person{
		Name: Name{Family: "Newmarch", Personal: "Server"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	service := "0.0.0.0:1200" // localhost:1200

	// get *IPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	// get *TCPListener
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)


	for {

		// get connection for client, do nothing if cannot get connection
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// create json encoder for connection
		encoder := json.NewEncoder(conn)

		// create json decoder for connection
		decoder := json.NewDecoder(conn)

		//
		for n := 0; n < 1; n++ {

			// decode data from client
			var personClient Person
			decoder.Decode(&personClient)
			fmt.Println(personClient.String())

			// encode data back to sent as response to client
			encoder.Encode(personServer)
		}

		conn.Close() // we're finished
	}
}

// checkError print Fatal error to stdout and quit execution if err not nil
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}