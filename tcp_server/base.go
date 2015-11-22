package tcp_server

func Init() {

	server := New(":9999")
	server.Listen()
}
