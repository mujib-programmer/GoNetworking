// UnmarshalXML
// unmarshal xml string to go struct data type and then print their data to stdout
//
// You can execute the UnmarshalXML like this:
//		UnmarshalXML
//
// Go 1 also has support for marshalling data structures into an XML document. 
// The function is:
//		func Marshal(v interface}{) ([]byte, error)

package main

import (
	"encoding/xml"
	"fmt"
	"os"
	//"strings"
)

type Person struct {
	XMLName Name `xml:"person"`
	Name Name `xml:"name"`
	Email []Email `xml:"email"`
}

type Name struct {
	Family string `xml:"family"`
	Personal string `xml:"personal"`
}

type Email struct {
	Type string `xml:"type,attr"`
	Address string `xml:",chardata"`
}

func main() {
	
	// xml string
	str := `<?xml version="1.0" encoding="utf-8"?>
	<person>
	<name>
	<family> Newmarch </family>
	<personal> Jan </personal>
	</name>
	<email type="personal">
	jan@newmarch.name
	</email>
	<email type="work">
	j.newmarch@boxhill.edu.au
	</email>
	</person>`
	
	var person Person
	
	// unmarshal byte slice str to person struct
	err := xml.Unmarshal([]byte(str), &person)
	checkError(err)
	
	// now use the person structure e.g.
	fmt.Println("Family name: \"" + person.Name.Family + "\"")
	fmt.Println("Second email address: \"" + person.Email[1].Address + "\"")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}