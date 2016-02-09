package tcp_server

import (
	"net"
	"testing"
	"time"
)

func testServer() *server {
	return New("localhost:9999")
}

func TestAcceptingNewClient(t *testing.T) {
	server := testServer()

	var newClient bool

	server.OnNewClientCallback(func(conn net.Conn) {
		newClient = true
	})
	go server.Listen()

	time.Sleep(10 * time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		t.Fatal("Failed to connect to test server")
	}

	conn.Close()


	time.Sleep(10 * time.Millisecond)

	if !newClient {
		t.Fail()
	}
}
