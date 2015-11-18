// SaveJSON
// Save json data to json file
//
// You can execute the SaveJSON like this:
// 		SaveJSON
//
// you will get person.json file in current directory as result

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

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	saveJSON("person.json", person)
}

func saveJSON(fileName string, key interface{}) {

	// create json file
	outFile, err := os.Create(fileName)
	checkError(err)

	// create encoder for json file
	encoder := json.NewEncoder(outFile)

	// encode key/data
	err = encoder.Encode(key)
	checkError(err)

	// close file
	outFile.Close()
}

// checkError print Fatal error to stdout and quit execution if err not nil
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}