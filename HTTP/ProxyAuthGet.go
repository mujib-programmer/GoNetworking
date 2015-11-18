// ProxyAuthGet
// Some proxies will require authentication, by a user name and password in order 
// to pass requests. A common scheme is "basic authentication" in which the user name
// and password are concatenated into a string "user:password" and then BASE64 encoded.
// This is then given to the proxy by the HTTP request header "Proxy-Authorisation" 
// with the flag that it is the basic authentication
//
// The following program illlustrates this, adding the Proxy-Authentication header 
// to the previous ProxyGet program
//
// If you have a proxy at, say, XYZ.com on port 8080, 
// You can execute the ClientGet like this:
// 		ProxyGet http://XYZ.com:8080/ http://www.google.co

package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const auth = "jannewmarch:mypassword"

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ", os.Args[0], "http://proxy-host:port http://host:port/page")
		os.Exit(1)
	}
	
	proxy := os.Args[1]
	
	proxyURL, err := url.Parse(proxy)
	checkError(err)
	
	rawURL := os.Args[2]
	
	url, err := url.Parse(rawURL)
	checkError(err)
	
	// encode the auth
	basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	
	client := &http.Client{Transport: transport}
	
	request, err := http.NewRequest("GET", url.String(), nil)
	
	request.Header.Add("Proxy-Authorization", basic)
	
	dump, _ := httputil.DumpRequest(request, false)
	
	fmt.Println(string(dump))
	
	// send the request
	response, err := client.Do(request)
	checkError(err)
	
	fmt.Println("Read ok")
	
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	
	fmt.Println("Reponse ok")
	
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

func checkError(err error) {
	if err != nil {
		if err == io.EOF {
			return
		}
		
		fmt.Println("Fatal error ", err.Error())
		
		os.Exit(1)
	}
}
