package main

import (
	"log"
	"http"
	"os"
	"net"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	
	log.Stdout("Starting basic HTTP server...\n")
	
	cwd, cwdErr := os.Getwd()
	if cwdErr != nil {
		fatalError(cwdErr.String())
	}
	
	tcp_listen_addr := &net.TCPAddr{net.IPv4(0,0,0,0), 0}
	listener, listenErr := net.ListenTCP("tcp", tcp_listen_addr)
	
	if listenErr != nil {
		log.Stderr(listenErr.String())
	}
	log.Stdoutf("Listening on port: %s\n", listener.Addr().String())
	
	fsHandler := http.FileServer(cwd, "")
	http.Serve(listener, fsHandler)
	listener.Close()
	log.Stdout("Done.")
}

func fatalError(s string) {
	log.Stderrf("%s\n", s)
	os.Exit(1)
}
