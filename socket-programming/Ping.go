// Ping
// Ping uses the "echo" command from the ICMP protocol. This is a byte-oriented protocol,
// in which the client sends a stream of bytes to another host, and the host replies. the format is:
//		* The first byte is 8, standing for the echo message
//		* The second byte is zero
//		* The third and fourth bytes are a checksum on the entire message
//		* The fifth and sixth bytes are an arbitrary indentifier
//		* The seventh and eight bytes are an arbitrary sequence number
//		* The rest of the packet is user data
//
// The following program will prepare an IP connection, send a ping request to a host and get a reply.
// You may need to have root access in order to run it successfully.

package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	// Ping need 1 argument/parameter to execute
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}

	// get *IPAddr
	addr, err := net.ResolveIPAddr("ip", os.Args[1])
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}

	// get IPConn
	conn, err := net.DialIP("ip4:icmp", addr, addr)
	checkError(err)


	var msg [512]byte
	msg[0] = 8 // echo
	msg[1] = 0 // code 0
	msg[2] = 0 // checksum, fix later
	msg[3] = 0 // checksum, fix later
	msg[4] = 0 // identifier[0]
	msg[5] = 13 //identifier[1]
	msg[6] = 0 // sequence[0]
	msg[7] = 37 // sequence[1]

	len := 8
	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	// write ICMP data to send
	_, err = conn.Write(msg[0:len])
	checkError(err)

	// read response from server
	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("Got response")

	if msg[5] == 13 {
		fmt.Println("identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}

	os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	sum := 0
	// assume even for now
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}