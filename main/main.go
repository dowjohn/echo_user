package main

import (
	"log"
	"user/database"
	"user/echoserver"
	"user/env"
)

func main() {
	println("initializing users")
	config, err := env.Init()
	if err != nil {
		log.Fatal("failed to read config")
	}
	connection, err := database.Init(config)
	if err != nil {
		println("failed db connection")
		return
	}
	echoserver.Init(connection)
}