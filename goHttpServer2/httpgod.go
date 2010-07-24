package main

import (
	"log"
	"http"
	"os"
	"net"
	"runtime"
)

var (
	Resp []byte = []byte("Hello!\n")
)

func handler(conn *http.Conn, req *http.Request) {
	conn.SetHeader("Content-Type", "text/plain")
	conn.Write(Resp)
}

func main() {
	runtime.GOMAXPROCS(4)
	
	log.Stdout("Starting basic HTTP server...\n")
	
	tcp_listen_addr := &net.TCPAddr{net.IPv4(0,0,0,0), 0}
	listener, listenErr := net.ListenTCP("tcp", tcp_listen_addr)
	
	if listenErr != nil {
		log.Stderr(listenErr.String())
	}
	log.Stdoutf("Listening on port: %s\n", listener.Addr().String())
	
	http.HandleFunc("/hello", handler)
	http.Serve(listener, nil)
	listener.Close()
	log.Stdout("Done.")
}

func fatalError(s string) {
	log.Stderrf("%s\n", s)
	os.Exit(1)
}
