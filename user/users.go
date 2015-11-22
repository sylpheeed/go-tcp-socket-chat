package user
import (
	"net"
)

type users map[int]*User

var Users users = make(users, 0)
var counter int = 0

func Create(connection net.Conn) {
	counter += 1
	user := &User{
		Id: counter,
		connection: connection,
	}
	Users[counter] = user
	go user.listen()
	user.newUserConnected()
}

func (u users) Emit(message string) {
	for _, el := range Users {
		el.connection.Write([]byte(message + "\n"))
	}
}
