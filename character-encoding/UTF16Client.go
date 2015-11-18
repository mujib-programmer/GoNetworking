// UTF16Client
// a client that reads a byte stream, extracts and examines the BOM and then decodes the rest of the stream.
//
// You can execute the UTF16Client like this:
// 		UTF16Client 0.0.0.0:1200
//		or
//		UTF16Client localhost:1200
//
// Which will output something like:
//		Response from server: j'ai arrêté

package main
import (
	"fmt"
	"net"
	"os"
	"unicode/utf16"
)
const BOM = '\ufffe'
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)
	shorts := readShorts(conn)
	ints := utf16.Decode(shorts)
	str := string(ints)
	fmt.Println("Response from server:", str)
	os.Exit(0)
}
func readShorts(conn net.Conn) []uint16 {
	var buf [512]byte
	// read everything into the buffer
	n, err := conn.Read(buf[0:2])
	for true {
		m, err := conn.Read(buf[n:])
		if m == 0 || err != nil {
			break
		}
		n += m
	}
	checkError(err)
	var shorts []uint16
	shorts = make([]uint16, n/2)
	if buf[0] == 0xff && buf[1] == 0xfe {
		// big endian
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buf[i])<<8 + uint16(buf[i+1])
		}
	} else if buf[1] == 0xff && buf[0] == 0xfe {
		// little endian
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buf[i+1])<<8 + uint16(buf[i])
		}
	} else {
		// unknown byte order
		fmt.Println("Unknown order")
	}
	return shorts
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