package service

import (
	"github.com/Manifoldz/TodoListRESTAPI/internal/entities"
	"github.com/Manifoldz/TodoListRESTAPI/pkg/repository"
)

type TodoItemService struct {
	repo     repository.ToDoItem
	listRepo repository.ToDoList
}

func NewTodoItemService(repo repository.ToDoItem, listRepo repository.ToDoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (s *TodoItemService) Create(listId int, item entities.ToDoItem) (int, error) {
	_, err := s.listRepo.GetById(listId)
	if err != nil {
		// list not found
		return 0, err
	}
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(listId int) ([]entities.ToDoItem, error) {
	return s.repo.GetAll(listId)
}
