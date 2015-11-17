// MD5Hash
// Using MD5 to hashing "hello\n" string
//
// You can execute the FTPServer like this:
// 		MD5Hash
//
// Which will output something like:
//		b1946ac9 2492d234 7c6235b4 d2611184

package main
import (
	"crypto/md5"
	"fmt"
)
func main() {
	hash := md5.New()

	bytes := []byte("hello\n")

	hash.Write(bytes)

	hashValue := hash.Sum(nil)

	hashSize := hash.Size()

	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n])<<24 +
		uint32(hashValue[n+1])<<16 +
		uint32(hashValue[n+2])<<8 +
		uint32(hashValue[n+3])
		fmt.Printf("%x ", val)
	}
	fmt.Println()
}