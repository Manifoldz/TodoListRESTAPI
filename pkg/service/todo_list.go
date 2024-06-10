package service

import (
	"github.com/Manifoldz/TodoListRESTAPI/internal/entities"
	"github.com/Manifoldz/TodoListRESTAPI/pkg/repository"
)

type TodoListService struct {
	repo repository.ToDoList
}

func NewTodoListService(repo repository.ToDoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(list entities.ToDoList) (int, error) {
	return s.repo.Create(list)
}

func (s *TodoListService) GetAll() ([]entities.ToDoList, error) {
	return s.repo.GetAll()
}
func (s *TodoListService) GetById(id int) (entities.ToDoList, error) {
	return s.repo.GetById(id)
}
