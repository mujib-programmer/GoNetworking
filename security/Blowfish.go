// Blowfish
// Using blowfish algorithm to encrypt "hello\n\n\n" string.
// Blowfish is not in the Go 1 distribution. Instead it is on the http://code.google.com/p/ site.
// You have to install it by running "go get code.google.com/p/go.crypto/blowfish"
// in a directory where you have source that needs to use it.
//
// You can execute the FTPServer like this:
// 		Blowfish
//
// Which will output something like:
//		hello
//
//
//
//

package main

import (
	"bytes"
	"code.google.com/p/go.crypto/blowfish"
	"fmt"
)

func main() {
	key := []byte("my key")

	cipher, err := blowfish.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	src := []byte("hello\n\n\n")

	var enc [512]byte

	cipher.Encrypt(enc[0:], src)

	var decrypt [8]byte

	cipher.Decrypt(decrypt[0:], enc[0:])

	result := bytes.NewBuffer(nil)

	result.Write(decrypt[0:8])

	fmt.Println(string(result.Bytes()))
}