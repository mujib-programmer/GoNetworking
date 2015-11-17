// JSONEchoClient
// encode json data to send to server, and then decode marshalled response data from server
//
// You can execute the JSONEchoClient like this:
// 		JSONEchoClient localhost:1200
//
// Which will output something like this
//      Server Newmarch
//		home: jan@newmarch.name
//		work: j.newmarch@boxhill.edu.au

package main

import (
	"fmt"
	"net"
	"os"
	"encoding/json"
	"bytes"
	"io"
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
	personClient := Person{
		Name: Name{Family: "Newmarch", Personal: "Client"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	// JSONEchoClient need 1 argument/parameter (host:port)
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}

	service := os.Args[1] // host:port

	// create connection to host:port
	conn, err := net.Dial("tcp", service)
	checkError(err)

	// create json encoder for connection
	encoder := json.NewEncoder(conn)

	// create json decoder for connection
	decoder := json.NewDecoder(conn)


	for n := 0; n < 1; n++ {

		// encode person data to send to server
		encoder.Encode(personClient)

		// decode person data from server
		var personServer Person
		decoder.Decode(&personServer)

		fmt.Println(personServer.String())
	}

	os.Exit(0)
}

// checkError print Fatal error to stdout and quit execution if err not nil
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

// implement readFully method for net.Conn struct
func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
