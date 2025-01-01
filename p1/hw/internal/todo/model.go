// internal/todo/model.go
package todo

import "time"

type Task struct {
	ID          string
	Title       string
	Description string
	CreateadAt  time.Time
	UpdatedAt   time.Time
	Done        bool
}

// Task - структура
// с помощью которой будет передаваться и храниться информация о задаче
// из запроса со стороны клиента
