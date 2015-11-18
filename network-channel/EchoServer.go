// EchoServer
// In order to make a channel visible to clients, you need to export it. This is done by creating an exporter using NewExporter with no
// parameters. The server then calls ListenAndServe to lsiten and handle responses. This takes two parameters, the first being the
// underlying transport mechanism such as "tcp" and the second being the network listening address (usually just a port number.
//
// For each channel, the server creates a normal local channel and then calls Export to bind this to the network channel. At the time of
// export, the direction of communication must be specified. Clients search for channels by name, which is a string. This is specified to
//the exporter.
//
// The server then uses the local channels in the normal way, reading or writing on them. We illustrate with an "echo" server which
// reads lines and sends them back. It needs two channels for this. The channel that the client writes to we name "echo-out". On the
// server side this is a read channel. Similarly, the channel that the client reads from we call "echo-in", which is a write channel to the
// server.

package main
import (
	"fmt"
	"os"
    "golang.org/x/exp/old/netchan"
)
func main() {

	exporter := netchan.NewExporter()

	err := exporter.ListenAndServe("tcp", ":2345")
	checkError(err)

	// create channel echoIn and echoOut
	echoIn := make(chan string)
	echoOut := make(chan string)

	exporter.Export("echo-in", echoIn, netchan.Send)
	exporter.Export("echo-out", echoOut, netchan.Recv)

	for {
		fmt.Println("Getting from echoOut")

		s, ok := <-echoOut
		if !ok {
			fmt.Printf("Read from channel failed")
			os.Exit(1)
		}

		fmt.Println("received", s)

		fmt.Println("Sending back to echoIn")

		echoIn <- s
		fmt.Println("Sent to echoIn")
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}