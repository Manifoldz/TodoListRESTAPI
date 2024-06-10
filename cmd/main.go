package main

import (
	"log"

	"github.com/Manifoldz/TodoListRESTAPI/internal/server"
	"github.com/Manifoldz/TodoListRESTAPI/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	svr := new(server.Server)
	if err := svr.Start("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while starting http server: %s", err.Error())
	}
}
