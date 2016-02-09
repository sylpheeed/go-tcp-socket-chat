package main
import (
	"github.com/sylpheeed/go-tcp-socket-chat"
	"github.com/sylpheeed/go-tcp-socket-chat/examples/user"
	"net"
)

func main() {
	server := tcp_server.New(":9999")
	server.OnNewClientCallback(func(conn net.Conn) {
		user.Create(conn)
	})
	server.Listen()
}

