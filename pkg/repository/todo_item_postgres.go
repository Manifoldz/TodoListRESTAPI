package repository

import (
	"fmt"

	"github.com/Manifoldz/TodoListRESTAPI/internal/entities"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item entities.ToDoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id;", toDoItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES  ($1, $2);", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(listId int) ([]entities.ToDoItem, error) {
	var items []entities.ToDoItem
	query := fmt.Sprintf("SELECT t1.id, t1.title, t1.description, t1.done FROM %s t1 INNER JOIN %s t2 ON t2.item_id  = t1.id WHERE t2.list_id  =  $1;", toDoItemsTable, listsItemsTable)

	if err := r.db.Select(&items, query, listId); err != nil {
		return nil, err
	}

	return items, nil
}
