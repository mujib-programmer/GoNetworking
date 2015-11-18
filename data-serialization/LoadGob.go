// LoadGob
// load person.gob data and decoded with gob
//
// You can execute the LoadGob like this:
// 		LoadGob
//
// which will output something like this
//		Person Jan Newmarch
//		home: jan@newmarch.name
//		work: j.newmarch@boxhill.edu.au
//
// Gob is a serialisation technique specific to Go. It is designed to encode Go data types specifically and does not at present have
// support for or by any other languages. It supports all Go data types except for channels, functions and interfaces. It supports
// integers of all types and sizes, strings and booleans, structs, arrays and slices. At present it has some problems with circular
// structures such as rings, but that will improve over time.

package main

import (
	"fmt"
	"os"
	"encoding/gob"
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

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

func main() {
	var person Person
	loadGob("person.gob", &person)
	fmt.Println("Person", person.String())
}

func loadGob(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}

// checkError print Fatal error to stdout and quit execution if err not nil
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}