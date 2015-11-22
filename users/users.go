package users
import (
	"net"
	"bufio"
)

type User struct {
	Name       string
	Id         int
	connection net.Conn
}


type allUsers map[int]*User

var AllUsers allUsers = make(allUsers, 0)
var counter int = 0

func Create(connection net.Conn) {
	counter += 1
	user := &User{
		Id: counter,
		connection: connection,
	}
	AllUsers[counter] = user
	go user.listen()
	user.newUserConnected()
}

func (u *User) Broadcast(message string) {
	for id, el := range AllUsers {
		if u.Id != id {
			el.connection.Write([]byte("[" + u.Name + "]: " + message + "\n"))
		}
	}
}

func (u *User) Emit(message string) {
	u.connection.Write([]byte(message + "\n"))
}

func (u *User) Quit() {
	u.connection.Close()
	delete(AllUsers, u.Id)
}

// Read client data from channel
func (u *User) listen() {
	reader := bufio.NewReader(u.connection)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			u.Quit()
		}else {
			u.newMessage(message)
		}
	}
}

func (u *User) newMessage(message string) {
	if u.Name == "" {
		u.Name = message
		u.Broadcast("New user " + u.Name + " is connected to chat")
	}else {
		u.Broadcast(message)
	}
}

func (u *User) newUserConnected() {
	u.Emit("Tell me your name.")
}

func (u *User) userConnectionClose() {
	u.Broadcast("User " + u.Name + " has left the chat")
}
