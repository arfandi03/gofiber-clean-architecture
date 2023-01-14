package entity

import (
	"golang-todo-app/enum"

	"github.com/google/uuid"
)

type Todo struct {
	Id          uuid.UUID       `gorm:"primaryKey;column:todo_id;type:varchar(36)"`
	Title       string          `gorm:"column:title"`
	Description string          `gorm:"column:description"`
	Status      enum.TodoStatus `gorm:"type:enum('OPEN','WIP','COMPLETED');column:status"`
	Username    string          `gorm:"column:username"`
}
