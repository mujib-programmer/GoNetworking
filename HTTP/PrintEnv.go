// PrintEnv
// serve all files inside /var/www/ directory on the server if we use "/" url pattern.
// Or print the values of environment variables if client use "/cgi-bin/printenv" url pattern.
//
// You can execute the PrintEnv like this:
//		PrintEnv
//
// You can test it by creating several files on /var/www/ directory on server,
// and then paste this url "http://localhost:8000/" on browser.
//
// If you typed this url on your web browser: "http://localhost:8000/cgi-bin/printenv"
// you will get list of the values of environment variables
// 

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	
	// file handler for most files
	fileServer := http.FileServer(http.Dir("/var/www"))
	
	http.Handle("/", fileServer)
	
	// function handler for /cgi-bin/printenv
	http.HandleFunc("/cgi-bin/printenv", printEnv)
	
	// deliver requests to the handlers
	err := http.ListenAndServe(":8000", nil)
	
	checkError(err)
	// That's it!
}
	
func printEnv(writer http.ResponseWriter, req *http.Request) {
	
	env := os.Environ()
	
	writer.Write([]byte("<h1>Environment</h1>\n<pre>"))
	
	for _, v := range env {
		writer.Write([]byte(v + "\n"))
	}
	
	writer.Write([]byte("</pre>"))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}