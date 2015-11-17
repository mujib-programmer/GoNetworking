// Base64Encoding
//

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {
	eightBitData := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	bb := &bytes.Buffer{}

	encoder := base64.NewEncoder(base64.StdEncoding, bb)

	encoder.Write(eightBitData)

	encoder.Close()

	fmt.Println("encoded eghtBitData: ", bb)

	dbuf := make([]byte, 10)

	decoder := base64.NewDecoder(base64.StdEncoding, bb)

	decoder.Read(dbuf)

	fmt.Print("Decoded bb: ")
	for _, ch := range dbuf {
		fmt.Print(ch)
	}
	fmt.Println()

}