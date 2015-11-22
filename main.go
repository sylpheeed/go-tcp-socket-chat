package main
import (
	"github.com/sylpheeed/go-tcp-socket-chat/tcp_server"
	"github.com/sylpheeed/go-tcp-socket-chat/users"
	"net"
)

func main() {
	server := tcp_server.New(":9999")
	server.Listen()
	server.OnNewClientCallback(func(conn net.Conn) {
		users.Create(conn)
	})
}

