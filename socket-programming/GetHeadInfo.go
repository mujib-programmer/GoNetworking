// GetHeadInfo
// the program (GetHeadInfo.go) to establish the connection for a TCP address, send the request string, read and print the response.
//
// Once compiled it can be invoked by e.g.
//     GetHeadInfo www.google.com:80
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
	"net"
	"os"
	"fmt"
	"io/ioutil"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	// create a TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// get a TCPConn for communication.
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	// send message to server
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)


	// result, err := readFully(conn)

	// get responses from server
	result, err := ioutil.ReadAll(conn)
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
