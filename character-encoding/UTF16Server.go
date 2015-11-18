// UTF16Server
// Using the BOM convention, we can write a server that prepends a BOM and writes a string in UTF-16
//
// You can execute the UTF16Server like this:
// 		UTF16Server
//
// Which will output something like:
//		Waiting for client connection on 0.0.0.0:1200
//
// After some client connected to this server, server will print to their stdout something like:
//		Accept conncetion from  127.0.0.1:1200

package main
import (
	"fmt"
	"net"
	"os"
	"unicode/utf16"
)
const BOM = '\ufffe'
func main() {
	port:= ":1200"
	service := "0.0.0.0" + port // localhost:1200

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	fmt.Println("Waiting for client connection on", service)

	for {
		conn, err := listener.Accept()

		fmt.Println("Accept conncetion from ", conn.LocalAddr())

		if err != nil {
			continue
		}
		str := "j'ai arrêté"
		shorts := utf16.Encode([]rune(str))
		writeShorts(conn, shorts)
		conn.Close() // we're finished
	}
}
func writeShorts(conn net.Conn, shorts []uint16) {
	var bytes [2]byte
	// send the BOM as first two bytes
	bytes[0] = BOM >> 8
	bytes[1] = BOM & 255
	_, err := conn.Write(bytes[0:])
	if err != nil {
		return
	}
	for _, v := range shorts {
		bytes[0] = byte(v >> 8)
		bytes[1] = byte(v & 255)
		_, err = conn.Write(bytes[0:])
		if err != nil {
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


/*
UTF-16 and Go
-------------

UTF-16 deals with arrays of short 16-bit unsigned integers. The package utf16 is designed to manage such arrays. To convert a
normal Go string, that is a UTF-8 string, into UTF-16, you first extract the code points by coercing it into a []rune and then use
utf16.Encode to produce an array of type uint16 .
Similarly, to decode an array of unsigned short UTF-16 values into a Go string, you use utf16.Decode to convert it into code points
as type []rune and then to a string. The following code fragment illustrates this

	str := "百度一下,你就知道 "
	runes := utf16.Encode([]rune(str))
	ints := utf16.Decode(runes)
	str = string(ints)

These type conversions need to be applied by clients or servers as appropriate, to read and write 16-bit short integers, as shown
below.

Little-endian and big-endian
----------------------------

Unfortunately, there is a little devil lurking behind UTF-16. It is basically an encoding of characters into 16-bit short integers. The
big question is: for each short, how is it written as two bytes? The top one first, or the top one second? Either way is fine, as long
as the receiver uses the same convention as the sender.

Unicode has addressed this with a special character known as the BOM (byte order marker). This is a zero-width non-printing
character, so you never see it in text. But its value 0xfffe is chosen so that you can tell the byte-order:

	In a big-endian system it is FF FE
	In a little-endian system it is FE FF

Text will sometimes place the BOM as the first character in the text. The reader can then examine these two bytes to determine
what endian-ness has been used.
*/
