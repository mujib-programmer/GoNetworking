// GetHttp
//
// You can execute the GetHttp like this:
// 		GetHttp http://www.golang.com/
//
// Which will output something like:
//		HTTP/1.1 200 OK
//		Content-Length: 7839
//		Alt-Svc: quic=":443"; p="1"; ma=604800
//		Alternate-Protocol: 443:quic,p=1
//		Content-Type: text/html; charset=utf-8
//		Date: Tue, 17 Nov 2015 23:58:53 GMT
//		Server: Google Frontend

package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "url")
		os.Exit(1)
	}

	url := os.Args[1]

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	b, _ := httputil.DumpResponse(response, false)
	fmt.Print(string(b))

	contentTypes := response.Header["Content-Type"]

	if !acceptableCharset(contentTypes) {
		fmt.Println("Cannot handle", contentTypes)
		os.Exit(4)
	}

	var buf [512]byte
	reader := response.Body

	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}

	os.Exit(0)
}

func acceptableCharset(contentTypes []string) bool {
	// each type is like [text/html; charset=UTF-8]
	// we want the UTF-8 only
	for _, cType := range contentTypes {
		if strings.Index(cType, "UTF-8") != -1 {
			return true
		}
	}
	return false
}