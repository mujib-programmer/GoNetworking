// IPGetHeadInfo
// the program (IPGetHeadInfo.go) to establish the connection for a TCP/UDP address, send the request string, read and print the response.
//
// Once compiled it can be invoked by e.g.
//     IPGetHeadInfo www.google.com:80
//
// Which will output something like this
//     HTTP/1.0 302 Found
//     Location: http://www.google.co.id/?gws_rd=cr&ei=H8dJVt2OG8GPuASewKXIBg
//	   Cache-Control: private
//	   Content-Type: text/html; charset=UTF-8
//	   P3P: CP="This is not a P3P policy! See http://www.google.com/support/accounts/bin/answer.py?hl=en&answer=151657 for more info."
//	   Date: Mon, 16 Nov 2015 12:07:59 GMT
//	   Server: gws
//	   Content-Length: 261
//	   X-XSS-Protection: 1; mode=block
//	   X-Frame-Options: SAMEORIGIN
//	   Set-Cookie: PREF=ID=1111111111111111:FF=0:TM=1447675679:LM=1447675679:V=1:S=R2x2Er6b6OHXKAi-; expires=Thu, 31-Dec-2015 16:02:17 GMT; path=/; domain=.google.com
//	   Set-Cookie: NID=73=ZnaCVm2cqLwyKoHgqs1VRedu3YPFpKNOKcIQ6DcMEgk67wf5khTul-GpMpKmj5VtekXZDzxUGFBxa1n9G2FDbIvOIMnmKXZ7x8jaES8p7AzWtvvm4D3zggYg2h8TUPKKiFBKMs66Y6Gl_zXDLMHvExLkmR93s67syf4i; expires=Tue, 17-May-2016 12:07:59 GMT; path=/; domain=.google.com; HttpOnly

package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	// IPGetHeadInfo need 1 argument/parameter to execute
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		fmt.Println()
		os.Exit(1)
	}

	// accept first parameter (ip-address:port string) as service
	service := os.Args[1]

	// net.Dial method accept The net can be any of "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only),
	// "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only) and "ip6" IPv6-only) as first parameter
	conn, err := net.Dial("tcp", service)
	checkError(err)

	// send message to server
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// get responses from server
	result, err := readFully(conn)
	checkError(err)

	// print response to stdout
	fmt.Println(string(result))

	os.Exit(0)
}

// checkError print Fatal error to stderr and quit execution if err not nil
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// readFully read data accepted from client and returning result Byte buffer
func readFully(conn net.Conn) ([]byte, error) {

	// close connection on exit
	defer conn.Close()

	// create empty result bytes buffer
	result := bytes.NewBuffer(nil)

	// slice buf to read data from client
	var buf [512]byte

	for {

		// read data from client
		n, err := conn.Read(buf[0:])

		// write data from client to result buffer
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	// return result byte buffer
	return result.Bytes(), nil
}