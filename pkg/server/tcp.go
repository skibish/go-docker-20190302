package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/skibish/go-docker-20190302/pkg/commands"
)

// TCPServer is a server for a chat communication
type TCPServer struct {
	port            string
	l               net.Listener
	allClients      map[net.Conn]string
	newConnections  chan net.Conn
	deadConnections chan net.Conn
	messages        chan string
}

// NewTCPServer returns TCPServer
func NewTCPServer(port string) *TCPServer {
	return &TCPServer{
		port:            port,
		allClients:      make(map[net.Conn]string),
		newConnections:  make(chan net.Conn),
		deadConnections: make(chan net.Conn),
		messages:        make(chan string),
	}
}

// Start starts TCP server
func (srv *TCPServer) Start() error {
	l, err := net.Listen("tcp", ":"+srv.port)
	if err != nil {
		return err
	}

	srv.l = l

	go srv.engine()
	return nil
}

// Close closes server
func (srv *TCPServer) Close() error {
	return srv.Close()
}

func (srv *TCPServer) acceptConnections() {
	for {
		c, errAcc := srv.l.Accept()
		srv.allClients[c] = ""
		if errAcc != nil {
			log.Fatalf("failed to accept connection: %v\n", errAcc)
		}

		srv.newConnections <- c
	}
}

func (srv *TCPServer) readMessages(c net.Conn) {
	r := bufio.NewReader(c)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			log.Printf("failed to read: %v\n", err)
			break
		}

		srv.commands(c, s)
	}
}

func (srv *TCPServer) commands(c net.Conn, raw string) {
	split := strings.Split(raw, " ")
	cmd := split[0]

	switch cmd {
	case commands.Name:
		text := raw[len(cmd)+1 : len(raw)-1]
		srv.allClients[c] = text
		srv.messages <- fmt.Sprintf("[%s] CONNECTED\n", text)
	case commands.Message:
		text := raw[len(cmd)+1 : len(raw)-1]
		srv.messages <- fmt.Sprintf(">> [%s] %s\n", srv.allClients[c], text)
	case commands.Exit:
		srv.messages <- fmt.Sprintf("<< [%s] DISCONNECTED\n", srv.allClients[c])
		srv.deadConnections <- c
	}
}

func (srv *TCPServer) broadcast(m string) {
	for c := range srv.allClients {
		_, errWrite := c.Write([]byte(m))
		if errWrite != nil {
			log.Printf("failed to write: %v\n", errWrite)
			srv.deadConnections <- c
		}
	}
}

func (srv *TCPServer) delConnection(c net.Conn) {
	c.Close()
	delete(srv.allClients, c)
	log.Printf("client disconnected")
}

func (srv *TCPServer) engine() {
	go srv.acceptConnections()

	for {
		select {
		case c := <-srv.newConnections:
			go srv.readMessages(c)
		case m := <-srv.messages:
			log.Println("broadcasting meesage")
			go srv.broadcast(m)
		case c := <-srv.deadConnections:
			log.Println("deleting connection")
			go srv.delConnection(c)
		}
	}
}
