// ServerHandler
// simply returns a "204 No content" for all requests.
//
// You can execute the ServerHandler like this:
//		ServerHandler
// 
// HTTP requests received by a Go server are usually handled by a multiplexer 
// to examines the path in the HTTP request and calls the appropriate file handler, etc. 
// You can define your own handlers. These can either be registered with 
// the default multiplexer by calling http.HandleFunc which takes a pattern and a function. 
// The functions such as ListenAndServe then take a nil handler function. 
// This was done in the last example.
//
// If you want to take over the multiplexer role then you can give a non-zero function 
// as the handler function. This function will then be totally responsible 
// for managing the requests and responses.

package main

import (
	"net/http"
)

func main() {
	myHandler := http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		// Just return no content - arbitrary headers can be set, arbitrary body
		rw.WriteHeader(http.StatusNoContent)
	})

	http.ListenAndServe(":8000", myHandler)
}