package tcp_server

import (
//	"bufio"
	"log"
	"net"
	"github.com/sylpheeed/go-tcp-socket-chat/users"
)

// TCP server
type server struct {
	address string        // Address to open connection: localhost:9999
	joins   chan net.Conn // Channel for new connections
}

// Creates new User instance
func (s *server) newClient(conn net.Conn) {
	users.Create(conn)
}

// Listens new connections channel and creating new client
func (s *server) listenChannels() {
	for {
		select {
		case conn := <-s.joins:
			s.newClient(conn)
		}
	}
}

// Start network server
func (s *server) Listen() {
	go s.listenChannels()

	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal("Error starting TCP server.")
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		s.joins <- conn
	}
}

// Creates new tcp server instance
func New(address string) *server {
	log.Println("Creating server with address", address)
	server := &server{
		address: address,
		joins:   make(chan net.Conn),
	}

	return server
}
