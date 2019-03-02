package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to start tcp: %v\n", err)
	}
	defer l.Close()

	newConnections := make(chan net.Conn)
	deadConnetions := make(chan net.Conn)
	messages := make(chan string)

	allClients := make(map[net.Conn]bool)

	go func() {
		for {
			c, errAcc := l.Accept()
			if errAcc != nil {
				log.Fatalf("failed to accept connection: %v\n", errAcc)
			}

			newConnections <- c
		}
	}()

	for {
		select {
		case c := <-newConnections:
			log.Println("new connection")
			allClients[c] = true

			go func(c net.Conn) {
				r := bufio.NewReader(c)
				for {
					s, err := r.ReadString('\n')
					if err != nil {
						log.Printf("failed to read: %v\n", err)
						break
					}

					messages <- s
				}
			}(c)
		case message := <-messages:
			for c := range allClients {
				go func(c net.Conn, m string) {
					_, errWrite := c.Write([]byte(m))
					if errWrite != nil {
						log.Printf("failed to write: %v\n", errWrite)
						deadConnetions <- c
					}
				}(c, message)
			}
		case c := <-deadConnetions:
			c.Close()
			delete(allClients, c)
			log.Printf("client disconnected")
		}
	}
}
