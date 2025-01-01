// internal/todo/handler.go
package todo

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type TaskHandler struct {
	TaskRepository *TaskRepository
}

func NewTaskHandler(router *http.ServeMux, taskRepository *TaskRepository) {
	handler := &TaskHandler{
		TaskRepository: taskRepository,
	}
	router.HandleFunc("GET /tasks", handler.GetAll())
	router.HandleFunc("GET /tasks/{id}", handler.GetByID())
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

func (h *TaskHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		data, err := h.TaskRepository.GetByID(id)
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
		task.ID = uuid.New().String()
		task.CreateadAt = time.Now()
		task.UpdatedAt = time.Now()
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
		var taskDb *Task
		taskDb, err := h.TaskRepository.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		task.ID = id
		task.CreateadAt = taskDb.CreateadAt
		task.UpdatedAt = time.Now()
		_, err = h.TaskRepository.Update(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
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
