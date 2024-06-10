package entities

type ToDoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type ToDoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description" `
	Done        bool   `json:"done" `
}

type ListsItems struct {
	Id     int
	ListId int
	ItemId int
}
