package main

import "fmt" // Package implementing formatted I/O.
import "os"
import "net"
import "strings"

var handlers = 0

func main() {
	
	
	fmt.Printf("Starting Server That does nothing...\n")
	
	tcp_listen_addr := &net.TCPAddr{net.IPv4(0,0,0,0), 0}
	listener, error := net.ListenTCP("tcp", tcp_listen_addr)
	if error != nil {
		fmt.Printf(error.String())
	}
	fmt.Printf("Listening on port: %s\n", listener.Addr().String())
	
	for {
		go handle(listener.Accept())
	}
	
}

func handle(connection net.Conn, err os.Error) {
	handlers++
	defer func() {
		handlers--
	}()
	
	var handleId = handlers
	handleLog(handleId, "Starting")
	
	const readBufferSize = 512
	var readBuffer [readBufferSize]byte
	
	if err != nil {
		// Some error trying to listen
		handleLog(handleId, err.String())
		fatalError("Error when trying to handle connection\n")
	}
	
	for {
		nr, er := connection.Read(&readBuffer)
		switch {
			case nr < 0:
				fatalError(er.String());
			case nr == 0:  // EOF
				break
			case nr > 0:
		}

		sBuffer := string(readBuffer[0:nr])

		handleLog(handleId, sBuffer)

		if strings.HasSuffix(strings.TrimSpace(sBuffer), "end") {
			handleLog(handleId, "Received end command closing connection")
			break
		}
	}
	
	connection.Close()
}

func handleLog(h int, s string) {
	fmt.Printf("Handler %d >> %s\n", h, strings.TrimSpace(s))
}

func log(s string) {
	fmt.Printf("%s\n", s)
}

func fatalError(s string) {
	fmt.Fprintf(os.Stderr, "%s\n", s)
	os.Exit(1)
}
