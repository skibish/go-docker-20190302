package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(c net.Conn) {
	log.Println("New connection!")

	for {
		s, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Printf("failed to read string: %v\n", err)
			return
		}

		out := fmt.Sprintf("Hello, %s", s)

		_, errWrite := c.Write([]byte(out))
		if errWrite != nil {
			log.Printf("failed to write: %v\n", errWrite)
		}

	}
}

func main() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to start tcp: %v\n", err)
	}
	defer l.Close()

	for {
		c, errAcc := l.Accept()
		if errAcc != nil {
			log.Fatalf("failed to accept connection: %v\n", errAcc)
		}

		handleConnection(c)
	}

}
