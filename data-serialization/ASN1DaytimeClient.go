// ASN1DaytimeClient
// This connects to the service given in a form such as localhost:1200 , reads the TCP packet and decodes the ASN.1 content back
// into a string, which it prints.
//
// You can execute the ASN1DaytimeClient like this:
// 		ASN1DaytimeClient localhost:1200
//
// Which will output something like this
//      2015-11-17 06:30:48.690420607 +0700 WIB
//
// This client and server are exchanging ASN.1 encoded data values, not textual strings.

package main

import (
	"bytes"
	"encoding/asn1"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {

	// ASN1DaytimeClient need 1 argument/parameter to execute
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	// get host:port argument
	service := os.Args[1]

	// get conn
	conn, err := net.Dial("tcp", service)
	checkError(err)

	// result will contain all response data from server
	result, err := readFully(conn)
	checkError(err)

	var newtime time.Time

	// unmarshall data from server as time.Time type
	_, err1 := asn1.Unmarshal(result, &newtime)
	checkError(err1)

	// print response from server
	fmt.Println("After marshal/unmarshal: ", newtime.String())
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// readFully read all response data from server
func readFully(conn net.Conn) ([]byte, error) {

	// close connection on exit
	defer conn.Close()

	result := bytes.NewBuffer(nil)

	var buf [512]byte

	for {
		// read length data
		n, err := conn.Read(buf[0:])

		// copy data to result
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	// return result Bytes data
	return result.Bytes(), nil
}