// FTPServer
// waiting client to connect on port 1202
//
// You can execute the FTPServer like this:
// 		FTPServer
//
// this FTPServer only accept dir, cd, pwd, or quit command from client

package main

import (
	"fmt"
	"net"
	"os"
)

// list of accepted command from server to be processed by server
const (
	DIR = "DIR"
	CD = "CD"
	PWD = "PWD"
)

func main() {

	service := "0.0.0.0:1202" //localhost:1202

	// get *IPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	// get *TCPListener
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		// accept connection from client
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// handle client connection as goroutine/thread, so we can accept multiple client connection at a time
		go handleClient(conn)
	}
}

// handleClient
func handleClient(conn net.Conn) {

	// close connection on exit
	defer conn.Close()

	var buf [512]byte

	for {

		// read all data from client and store in buff
		n, err := conn.Read(buf[0:])
		if err != nil {
			conn.Close()
			return
		}

		// convert buf to string
		s := string(buf[0:n])

		// decode request
		if s[0:2] == CD {

			// execute change directory on server
			chdir(conn, s[3:])

		} else if s[0:3] == DIR {

			// get list of files and directory on current directory
			dirList(conn)

		} else if s[0:3] == PWD {

			// show current directory location
			pwd(conn)
		}
	}
}

// chdir execute change directory on server
func chdir(conn net.Conn, s string) {
	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

// pwd show current directory location
func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(s))
}

// dirList get list of files and directory on current directory
func dirList(conn net.Conn) {
	defer conn.Write([]byte("\r\n"))
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}
	for _, nm := range names {
		conn.Write([]byte(nm + "\r\n"))
	}
}

// checkError print Fatal error to stdout and quit execution if err not nil
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}