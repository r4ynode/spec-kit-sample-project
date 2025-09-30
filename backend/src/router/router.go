package router

import (
	"backend/src/infrastructure"
	"backend/src/usecase"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

// SetupRoutes configures the API routes.
func SetupRoutes(r *mux.Router, db *gorm.DB) {
	taskRepo := infrastructure.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskInteractor(taskRepo)
	taskHandler := NewTaskHandler(taskUseCase)

	r.Use(loggingMiddleware)

	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s %s in %v", r.Method, r.RequestURI, time.Since(start))
	})
}

type TaskHandler struct {
	taskUseCase usecase.TaskUseCase
}

func NewTaskHandler(taskUseCase usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{taskUseCase: taskUseCase}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.taskUseCase.GetAll()
	if err != nil {
		log.Printf("Error getting tasks: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var params usecase.CreateTaskParam
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Printf("Error decoding create task request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task, err := h.taskUseCase.Create(params)
	if err != nil {
		log.Printf("Error creating task: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error parsing task ID for update: %v", err)
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var params usecase.UpdateTaskParam
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Printf("Error decoding update task request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.taskUseCase.Update(id, params)
	if err != nil {
		log.Printf("Error updating task: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error parsing task ID for delete: %v", err)
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	if err := h.taskUseCase.Delete(id); err != nil {
		log.Printf("Error deleting task: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}