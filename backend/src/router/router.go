package router

import (
	"backend/src/infrastructure"
	"backend/src/usecase"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"encoding/json"
	"net/http"
	"strconv"
)

// SetupRoutes configures the API routes.
func SetupRoutes(r *mux.Router, db *gorm.DB) {
	taskRepo := infrastructure.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskInteractor(taskRepo)
	taskHandler := NewTaskHandler(taskUseCase)

	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var params usecase.CreateTaskParam
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task, err := h.taskUseCase.Create(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var params usecase.UpdateTaskParam
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.taskUseCase.Update(id, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	if err := h.taskUseCase.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}