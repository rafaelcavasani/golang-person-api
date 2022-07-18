package main

import "person-api/server"

func main() {
	server := server.NewServer()
	server.Run()
}
