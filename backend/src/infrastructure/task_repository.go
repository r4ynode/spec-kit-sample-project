package infrastructure

import (
	"gorm.io/gorm"

	"backend/src/entity"
)

// TaskRepository is a repository for tasks.
//
//counterfeiter:generate . TaskRepository
type TaskRepository struct {
	db *gorm.DB
}

// NewTaskRepository creates a new TaskRepository.
func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// Create creates a new task.
func (r *TaskRepository) Create(task *entity.Task) (*entity.Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

// FindAll returns all tasks.
func (r *TaskRepository) FindAll() ([]*entity.Task, error) {
	var tasks []*entity.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// FindByID returns a task by its ID.
func (r *TaskRepository) FindByID(id int) (*entity.Task, error) {
	var task entity.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// Update updates a task.
func (r *TaskRepository) Update(task *entity.Task) (*entity.Task, error) {
	if err := r.db.Save(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

// Delete deletes a task.
func (r *TaskRepository) Delete(task *entity.Task) error {
	return r.db.Delete(task).Error
}
