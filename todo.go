package todo

//TodoList ...
type TodoList struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Descrition string `json:"description"`
}

//UsersList ...
type UsersList struct {
	ID     int
	UserID int
	ListID int
}

//TodoItem ...
type TodoItem struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Descrition string `json:"description"`
	Done       bool   `json:"done"`
}

//ListsItems ...
type ListsItems struct {
	ID     int
	ItemID int
	ListID int
}
