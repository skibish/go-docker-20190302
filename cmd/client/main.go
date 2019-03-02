package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/skibish/go-docker-20190302/pkg/commands"
)

func main() {
	var (
		host = flag.String("addr", "127.0.0.1:9999", "TCP Server address")
	)
	flag.Parse()

	c, err := net.Dial("tcp", *host)
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}
	defer c.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Eneter name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read: %v\n", err)
	}

	fmt.Fprintf(c, commands.CreateNameCmd(name))

	connRead := bufio.NewReader(c)

	go func() {
		for {
			message, errRead := connRead.ReadString('\n')
			if errRead != nil {
				log.Fatalf("failed to read from server: %v\n", errRead)
			}

			fmt.Print(message)
		}
	}()

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("failed to read string: %v", err)
		}

		fmt.Fprintf(c, commands.CreateMessageCmd(text))
	}
}
