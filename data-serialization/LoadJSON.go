// LoadJSON
// Load person.json file and then decode the data to print data in some format
//
// You can execute the LoadJSON like this:
// 		LoadJSON
//
// Which will output something like this
//      Person Jan Newmarch
//		home: jan@newmarch.name
//		work: j.newmarch@boxhill.edu.au


package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	var person Person

	loadJSON("person.json", &person)

	fmt.Println("Person", person.String())
}

// loadJSON open json file, read the encoded data, decode the encoded json data then close file.
func loadJSON(fileName string, key interface{}) {

	// open json file
	inFile, err := os.Open(fileName)
	checkError(err)

	// decoder for json file
	decoder := json.NewDecoder(inFile)

	// decode json data
	err = decoder.Decode(key)
	checkError(err)

	// close file
	inFile.Close()
}

// checkError print Fatal error to stdout and quit execution if err not nil
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}