// HeadHttp
//
// You can execute the HeadHttp like this:
// 		go run HeadHttp.go http://www.golang.com/
//
// Which will output something like:
//		200 OK
//		Content-Type: [text/html; charset=utf-8]
//		Date: [Tue, 17 Nov 2015 23:53:10 GMT]
//		Server: [Google Frontend]
//		Content-Length: [7839]
//		Alternate-Protocol: [443:quic,p=1]
//		Alt-Svc: [quic=":443"; p="1"; ma=604800]


package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "url")
		os.Exit(1)
	}

	url := os.Args[1]

	response, err := http.Head(url)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Println(response.Status)

	for k, v := range response.Header {
		fmt.Println(k+":", v)
	}

	os.Exit(0)
}