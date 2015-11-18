// FTPClient
// connect to FTPServer on port 1202
//
// You can execute the FTPClient like this:
// 		FTPClient localhost
//
// you will get interactive console to type command for server.
// FTPServer only accept dir, cd, pwd, or quit command from client.
//		dir [enter] : will show list of files on current directory server
//		pwd			: will show full path current directory
//		cd dir-name	: will change current directory to dir-name
//		quit		: will quit FTPClient connection

package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
	"bytes"
)

// strings used by the user interface
const (
	uiDir = "dir"
	uiCd = "cd"
	uiPwd = "pwd"
	uiQuit = "quit"
)

// strings used across the network
const (
	DIR = "DIR"
	CD = "CD"
	PWD = "PWD"
)

func main() {

	// at least need one argument/parameter
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}

	// get first parameter as host
	host := os.Args[1]

	// create connection to server at port 1202
	conn, err := net.Dial("tcp", host+":1202")
	checkError(err)

	// read standard input
	reader := bufio.NewReader(os.Stdin)

	for {

		// read a line of standard input
		line, err := reader.ReadString('\n')

		// lose trailing whitespace
		line = strings.TrimRight(line, " \t\r\n")
		if err != nil {
			break
		}

		// split into command + arg
		strs := strings.SplitN(line, " ", 2)

		// decode user request
		switch strs[0] {

		case uiDir:
			dirRequest(conn)

		case uiCd:
			if len(strs) != 2 {
				fmt.Println("cd <dir>")
				continue
			}
			fmt.Println("CD \"", strs[1], "\"")
			cdRequest(conn, strs[1])

		case uiPwd:
			pwdRequest(conn)

		case uiQuit:
			conn.Close()
			os.Exit(0)

		default:
			fmt.Println("Unknown command")
		}
	}
}

// dirRequest request list of files from server current directory
func dirRequest(conn net.Conn) {
	conn.Write([]byte(DIR + " "))
	var buf [512]byte
	result := bytes.NewBuffer(nil)
	for {
		// read till we hit a blank line
		n, _ := conn.Read(buf[0:])
		result.Write(buf[0:n])
		length := result.Len()
		contents := result.Bytes()
		if string(contents[length-4:]) == "\r\n\r\n" {
			fmt.Println(string(contents[0 : length-4]))
			return
		}
	}
}

// cdRequest request change server current directory to specified dir string
func cdRequest(conn net.Conn, dir string) {
	conn.Write([]byte(CD + " " + dir))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	if s != "OK" {
		fmt.Println("Failed to change dir")
	}
}

// pwdRequest request server current directory
func pwdRequest(conn net.Conn) {
	conn.Write([]byte(PWD))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	fmt.Println("Current dir \"" + s + "\"")
}

// checkError print Fatal error to stdout and quit execution if err not nil
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}