package main

import (
	"log"

	"github.com/Manifoldz/TodoListRESTAPI/internal/server"
	"github.com/Manifoldz/TodoListRESTAPI/pkg/handler"
	"github.com/Manifoldz/TodoListRESTAPI/pkg/repository"
	"github.com/Manifoldz/TodoListRESTAPI/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	svr := new(server.Server)
	if err := svr.Start("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while starting http server: %s", err.Error())
	}
}
