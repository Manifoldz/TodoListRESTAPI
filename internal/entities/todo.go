package entities

type ToDoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
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
