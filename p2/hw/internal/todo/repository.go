// internal/todo/repository.go
package todo

import (
	"gorm.io/gorm/clause"
	"hw/pkg/db"
)

type TaskRepository struct {
	Database *db.Db
}

// TaskRepository - репозиторий для работы с задачами

func NewTaskRepository(database *db.Db) *TaskRepository {
	return &TaskRepository{
		Database: database,
	}
}

// NewTaskRepository создает новый экземпляр репозитория

func (repo *TaskRepository) Create(task *Task) (*Task, error) {
	result := repo.Database.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

// Create добавляет задачу в базу данных

func (repo *TaskRepository) GetById(id string) (*Task, error) {
	data := Task{}
	result := repo.Database.First(&data, "task_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

// GetByID возвращает задачу из базы данных по идентификатору

func (repo *TaskRepository) GetAll() ([]Task, error) {
	var tasks []Task
	result := repo.Database.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

// GetAll возвращает все задачи из базы данных

func (repo *TaskRepository) Update(task *Task) (*Task, error) {
	result := repo.Database.Clauses(clause.Returning{}).Where("task_id = ?", task.TaskID).Updates(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

// Update обновляет задачу в базе данных

func (repo *TaskRepository) Delete(id string) error {
	task, err := repo.GetById(id)
	if err != nil {
		return err
	}
	result := repo.Database.Delete(task)
	return result.Error
}

// Delete удаляет задачу из базы данных
