package server

import (
	"bufio"
	"log"
	"net"
)

// TCPServer is a server for a chat communication
type TCPServer struct {
	port            string
	l               net.Listener
	allClients      map[net.Conn]bool
	newConnections  chan net.Conn
	deadConnections chan net.Conn
	messages        chan string
}

// NewTCPServer returns TCPServer
func NewTCPServer(port string) *TCPServer {
	return &TCPServer{
		port:            port,
		allClients:      make(map[net.Conn]bool),
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
		srv.allClients[c] = true
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

		srv.messages <- s
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
