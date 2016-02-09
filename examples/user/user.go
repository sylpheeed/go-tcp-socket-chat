package user
import (
	"net"
	"bufio"
)

type User struct {
	Name       string
	Id         int
	connection net.Conn
}

func (u *User) Broadcast(message string) {
	for id, el := range Users {
		if u.Id != id {
			el.connection.Write([]byte(message + "\n"))
		}
	}
}

func (u *User) Emit(message string) {
	u.connection.Write([]byte(message + "\n"))
}

func (u *User) Quit() {
	u.connection.Close()
	delete(Users, u.Id)
	u.userConnectionClose()
}

// Read client data from channel
func (u *User) listen() {
	reader := bufio.NewReader(u.connection)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			u.Quit()
			return
		}else {
			message = message[:len(message) - 1]
			u.newMessage(message)
		}
	}
}

func (u *User) newMessage(message string) {
	if u.Name == "" {
		u.Name = message
		u.Broadcast("New user " + u.Name + " is connected to chat")
		u.Emit("Hi " + u.Name + " and welcome to chat!")
	}else {
		Users.Emit("[" + u.Name + "]: " + message)
	}
}

func (u *User) newUserConnected() {
	u.Emit("Tell me your name.")
}

func (u *User) userConnectionClose() {
	u.Broadcast("User " + u.Name + " has left the chat")
}
