package usecase

import (
	"backend/src/entity"
	"time"
)

// TaskInteractor is the use case for tasks.
//
//counterfeiter:generate . TaskInteractor
type TaskInteractor struct {
	taskRepository TaskRepository
}

// NewTaskInteractor creates a new TaskInteractor.
func NewTaskInteractor(taskRepository TaskRepository) *TaskInteractor {
	return &TaskInteractor{taskRepository: taskRepository}
}

// Create creates a new task.
func (i *TaskInteractor) Create(param CreateTaskParam) (*entity.Task, error) {
	task := &entity.Task{
		Title:     param.Title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	return i.taskRepository.Create(task)
}

// GetAll returns all tasks.
func (i *TaskInteractor) GetAll() ([]*entity.Task, error) {
	return i.taskRepository.FindAll()
}

// GetByID returns a task by its ID.
func (i *TaskInteractor) GetByID(id int) (*entity.Task, error) {
	return i.taskRepository.FindByID(id)
}

// Update updates a task.
func (i *TaskInteractor) Update(id int, param UpdateTaskParam) (*entity.Task, error) {
	task, err := i.taskRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	task.Title = param.Title
	task.Completed = param.Completed
	return i.taskRepository.Update(task)
}

// Delete deletes a task.
func (i *TaskInteractor) Delete(id int) error {
	task, err := i.taskRepository.FindByID(id)
	if err != nil {
		return err
	}
	return i.taskRepository.Delete(task)
}
