package usecase

import "backend/src/entity"

// TaskRepository is an interface for interacting with task data.
//
//counterfeiter:generate . TaskRepository
type TaskRepository interface {
	Create(task *entity.Task) (*entity.Task, error)
	FindAll() ([]*entity.Task, error)
	FindByID(id int) (*entity.Task, error)
	Update(task *entity.Task) (*entity.Task, error)
	Delete(task *entity.Task) error
}

// CreateTaskParam represents the parameters for creating a new task.
type CreateTaskParam struct {
	Title string `json:"title"`
}

// UpdateTaskParam represents the parameters for updating an existing task.
type UpdateTaskParam struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// TaskUseCase is an interface for task-related use cases.
//
//counterfeiter:generate . TaskUseCase
type TaskUseCase interface {
	Create(param CreateTaskParam) (*entity.Task, error)
	GetAll() ([]*entity.Task, error)
	GetByID(id int) (*entity.Task, error)
	Update(id int, param UpdateTaskParam) (*entity.Task, error)
	Delete(id int) error
}
