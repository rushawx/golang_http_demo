// internal/todo/handler.go
package todo

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type TaskHandler struct {
	TaskRepository *TaskRepository
}

func NewTaskHandler(router *http.ServeMux, taskRepository *TaskRepository) {
	handler := &TaskHandler{
		TaskRepository: taskRepository,
	}
	router.HandleFunc("GET /tasks", handler.GetAll())
	router.HandleFunc("GET /tasks/{id}", handler.GetById())
	router.HandleFunc("POST /tasks", handler.Create())
	router.HandleFunc("PUT /tasks/{id}", handler.Update())
	router.HandleFunc("DELETE /tasks/{id}", handler.Delete())
}

func (h *TaskHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := h.TaskRepository.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

func (h *TaskHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		data, err := h.TaskRepository.GetById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

func (h *TaskHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		task.Done = false
		data, err := h.TaskRepository.Create(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
	}
}

func (h *TaskHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		taskDb, err := h.TaskRepository.GetById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		var taskTitle string
		if task.Title != "" {
			taskTitle = task.Title
		} else {
			taskTitle = taskDb.Title
		}
		var taskDesc string
		if task.Description != "" {
			taskDesc = task.Description
		} else {
			taskDesc = taskDb.Description
		}
		var taskDone bool
		if task.Done {
			taskDone = true
		} else {
			taskDone = taskDb.Done
		}
		output, err := h.TaskRepository.Update(&Task{
			Model:       gorm.Model{ID: taskDb.Model.ID},
			TaskID:      taskDb.TaskID,
			Title:       taskTitle,
			Description: taskDesc,
			Done:        taskDone,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output)
	}
}

func (h *TaskHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		err := h.TaskRepository.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
