package repository

type ToDoList interface{}
type ToDoItem interface{}

type Repository struct {
	ToDoList
	ToDoItem
}

func NewRepository() *Repository {
	return &Repository{}
}
