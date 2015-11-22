package main
import (
	"github.com/sylpheeed/go-tcp-socket-chat/tcp_server"
	"github.com/sylpheeed/go-tcp-socket-chat/user"
	"net"
)

func main() {
	server := tcp_server.New(":9999")
	server.OnNewClientCallback(func(conn net.Conn) {
		user.Create(conn)
	})
	server.Listen()
}

