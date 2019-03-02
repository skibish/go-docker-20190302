package main

import (
	"flag"
	"log"

	"github.com/skibish/go-docker-20190302/pkg/server"
)

func main() {
	var (
		tcpPort = flag.String("tcp-port", "9999", "Port for TCP server")
	)
	flag.Parse()

	t := server.NewTCPServer(*tcpPort)
	err := t.Start()
	if err != nil {
		log.Fatalf("failed to start TCP server: %v", err)
	}
	defer t.Close()

	log.Println("application started!")
	select {}
}
