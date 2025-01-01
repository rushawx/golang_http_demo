// internal/todo/repository.go
package todo

import (
	"errors"
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
	repo.Database.Tasks[task.ID] = db.TaskDb{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CreateadAt:  task.CreateadAt,
		UpdatedAt:   task.UpdatedAt,
		Done:        task.Done,
	}
	return task, nil
}

// Create добавляет задачу в базу данных

func (repo *TaskRepository) GetByID(id string) (*Task, error) {
	task, ok := repo.Database.Tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	data := Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CreateadAt:  task.CreateadAt,
		UpdatedAt:   task.UpdatedAt,
		Done:        task.Done,
	}
	return &data, nil
}

// GetByID возвращает задачу из базы данных по идентификатору

func (repo *TaskRepository) GetAll() ([]Task, error) {
	var tasks []Task
	for _, task := range repo.Database.Tasks {
		tasks = append(tasks, Task{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			CreateadAt:  task.CreateadAt,
			UpdatedAt:   task.UpdatedAt,
			Done:        task.Done,
		})
	}
	return tasks, nil
}

// GetAll возвращает все задачи из базы данных

func (repo *TaskRepository) Update(task *Task) (*Task, error) {
	_, ok := repo.Database.Tasks[task.ID]
	if !ok {
		return nil, errors.New("task not found")
	}
	repo.Database.Tasks[task.ID] = db.TaskDb{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CreateadAt:  task.CreateadAt,
		UpdatedAt:   task.UpdatedAt,
		Done:        task.Done,
	}
	return task, nil
}

// Update обновляет задачу в базе данных

func (repo *TaskRepository) Delete(id string) error {
	_, ok := repo.Database.Tasks[id]
	if !ok {
		return errors.New("task not found")
	}
	delete(repo.Database.Tasks, id)
	return nil
}

// Delete удаляет задачу из базы данных
