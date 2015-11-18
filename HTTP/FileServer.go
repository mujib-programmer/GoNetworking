// FileServer
// serve all files inside /var/www/ directory on the server
//
// You can execute the FileServer like this:
//		FileServer
//
// After executed, FileServer will waiting for client to connect on port 8000.
// You can test it by creating several files on /var/www/ directory on server,
// and then paste this url http://localhost:8000/ on browser. Please note: FileServer
// and web browser (client) are assumed on the same machine (localhost).
//
// This server even delivers "404 not found" messages for requests for file resources 
// that don't exist!

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// deliver files from the directory /var/www
	fileServer := http.FileServer(http.Dir("/var/www/"))
	
	// register the handler and deliver requests to it
	err := http.ListenAndServe(":8000", fileServer)
	
	checkError(err)
	// That's it!
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}