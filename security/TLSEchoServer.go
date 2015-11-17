// TLSEchoServer
// Listening client connection using TLS (Transport Layer Security)
//
// You can execute the TLSEchoServer like this:
// 		TLSEchoServer
//
// after execution, program will create new files:
//		Listening
//
//	If you don't have a private.pem file for your private certificate, you can create one by executing another program:
//		GenRSAKeys
//
//	If you don't have a mujib.programmer.name.pem file for mujib.programmer certificate, you can create one by executing another program:
//		GenX509Cert

package main
import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"
)
func main() {
	cert, err := tls.LoadX509KeyPair("mujib.programmer.name.pem", "private.pem")
	checkError(err)
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	now := time.Now()
	config.Time = func() time.Time { return now }
	config.Rand = rand.Reader
	service := "0.0.0.0:1200"
	listener, err := tls.Listen("tcp", service, &config)
	checkError(err)
	fmt.Println("Listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("Accepted")
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		fmt.Println("Trying to read")
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}