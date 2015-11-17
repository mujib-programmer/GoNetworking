// SaveGob
// Save person data encoded with gob
//
// You can execute the SaveGob like this:
// 		SaveGob
//
// you will get person.gob file in current directory as result
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

func main() {

	// create person data
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	// encode person data with gob and save to person.gob file
	saveGob("person.gob", person)
}

// saveGob open file, encode person data, write to file, and then close file
func saveGob(fileName string, key interface{}) {

	// create file name to save encoded data
	outFile, err := os.Create(fileName)
	checkError(err)

	// create encoder
	encoder := gob.NewEncoder(outFile)

	// encode data on file fileName
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