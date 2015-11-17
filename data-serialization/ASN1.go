// ASN1

package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)

func main() {
	data := 13

	fmt.Println("Data: ", data)

	mdata, err := asn1.Marshal(data)
	checkError(err)

	fmt.Println("Marshaled data: ", mdata)

	var n int // var n to store unmarshaled mdata

	_, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)

	fmt.Println("Unmarshaled data: ", n)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
