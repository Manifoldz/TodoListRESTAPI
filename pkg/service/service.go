package service

import (
	"github.com/Manifoldz/TodoListRESTAPI/internal/entities"
	"github.com/Manifoldz/TodoListRESTAPI/pkg/repository"
)

type ToDoList interface {
	Create(list entities.ToDoList) (int, error)
}
type ToDoItem interface{}

type Service struct {
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		ToDoList: NewTodoListService(repos.ToDoList),
	}
}
