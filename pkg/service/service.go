package service

import "github.com/Manifoldz/TodoListRESTAPI/pkg/repository"

type ToDoList interface{}
type ToDoItem interface{}

type Service struct {
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
