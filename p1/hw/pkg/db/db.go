// pkg/db/db.go
package db

import "time"

type TaskDb struct {
	ID          string
	Title       string
	Description string
	CreateadAt  time.Time
	UpdatedAt   time.Time
	Done        bool
}

// Task - структура
// с помощью которой будет передаваться и храниться информация о задаче

type Db struct {
	Tasks map[string]TaskDb
}

// Db - база данных с хранением в оперативной памяти

func NewDb() *Db {
	return &Db{
		Tasks: map[string]TaskDb{},
	}
}

// NewDb создает новый экземпляр базы данных
