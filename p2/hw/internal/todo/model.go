// internal/todo/model.go
package todo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskID      uuid.UUID `json:"task_id" gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
}

// Task - структура
// с помощью которой будет передаваться и храниться информация о задаче
// из запроса со стороны клиента
