// LoadX509Cert
// Read certificate from mujib.programmer.name.cer file
//
// You can execute the FTPServer like this:
// 		LoadX509Cert
//
// Which will output something like:
//		Name mujib.programmer.name
//		Not before 2015-11-17 22:14:02 +0000 UTC
//		Not after 2016-11-16 22:14:02 +0000 UTC
//
//	If you don't have a mujib.programmer.name.cer file for your private key, you can create one by executing another program:
//		GenX509Cert


package main

import (
	"crypto/x509"
	"fmt"
	"os"
)

func main() {

	certCerFile, err := os.Open("mujib.programmer.name.cer")
	checkError(err)

	derBytes := make([]byte, 1000) // bigger than the file

	count, err := certCerFile.Read(derBytes)
	checkError(err)

	certCerFile.Close()

	// trim the bytes to actual length in call
	cert, err := x509.ParseCertificate(derBytes[0:count])
	checkError(err)

	fmt.Printf("Name %s\n", cert.Subject.CommonName)
	fmt.Printf("Not before %s\n", cert.NotBefore.String())
	fmt.Printf("Not after %s\n", cert.NotAfter.String())

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}