package main

import (
	"log"

	"github.com/Manifoldz/TodoListRESTAPI/internal/server"
)

func main() {
	svr := new(server.Server)
	if err := svr.Start(); err != nil {
		log.Fatalf("error while starting http server: %s", err.Error())
	}
}
