// GobEchoServer
// decode gob marshalled data from client, and then encode back data to sent as response to client
//
// You can execute the ASN1DaytimeClient like this:
// 		GobEchoServer
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
	"encoding/gob"
)
type Person struct {
	Name Name
	Email []Email
}
type Name struct {
	Family 	string
	Personal string
}
type Email struct {
	Kind string
	Address string
}
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

	service := "0.0.0.0:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)
		for n := 0; n < 1; n++ {
			var person Person
			decoder.Decode(&person)
			fmt.Println(person.String())

			encoder.Encode(personServer)
		}
		conn.Close() // we're finished Data serialisation
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}