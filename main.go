package main

func main() {
	server := NewApiServer(":9090")
	server.Run()
}
