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

			go func(c net.Conn) {
				r := bufio.NewReader(c)
				for {
					s, err := r.ReadString('\n')
					if err != nil {
						log.Printf("failed to read: %v\n", err)
						break
					}

					_, errWrite := c.Write([]byte(s))
					if errWrite != nil {
						log.Printf("failed to write: %v\n", errWrite)
					}
				}
			}(c)
		}
	}
}
