package main

import (
	"user/database"
	"user/echoserver"
)

func main() {
	println("initializing users")
	connection, err := database.Init()
	if err != nil {
		println("failed db connection")
		return
	}
	echoserver.Init(connection)
}