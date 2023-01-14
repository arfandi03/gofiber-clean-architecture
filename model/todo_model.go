package model

type TodoModel struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Username    string `json:"username"`
}

type CreateTodo struct {
	Title       string `json:"title" validate:"required,max=32"`
	Description string `json:"description" validate:"required"`
}

type UpdateTodoStatus struct {
	Status string `json:"status" validate:"oneof=OPEN WIP COMPLETED"`
}
