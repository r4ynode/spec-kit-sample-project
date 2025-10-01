package usecase

import (
	"testing"

	"backend/src/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTaskRepository is a mock of TaskRepository.
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(task *entity.Task) (*entity.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*entity.Task), args.Error(1)
}

func (m *MockTaskRepository) FindAll() ([]*entity.Task, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Task), args.Error(1)
}

func (m *MockTaskRepository) FindByID(id int) (*entity.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Task), args.Error(1)
}

func (m *MockTaskRepository) Update(task *entity.Task) (*entity.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*entity.Task), args.Error(1)
}

func (m *MockTaskRepository) Delete(task *entity.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func TestTaskInteractor_Create(t *testing.T) {
	taskRepo := new(MockTaskRepository)
	taskInteractor := NewTaskInteractor(taskRepo)

	task := &entity.Task{Title: "Test Task"}
	taskRepo.On("Create", task).Return(task, nil)

	createdTask, err := taskInteractor.Create(task)

	assert.NoError(t, err)
	assert.Equal(t, task, createdTask)
	taskRepo.AssertExpectations(t)
}

func TestTaskInteractor_FindAll(t *testing.T) {
	taskRepo := new(MockTaskRepository)
	taskInteractor := NewTaskInteractor(taskRepo)

	tasks := []*entity.Task{{ID: 1, Title: "Test Task"}}
	taskRepo.On("FindAll").Return(tasks, nil)

	foundTasks, err := taskInteractor.FindAll()

	assert.NoError(t, err)
	assert.Equal(t, tasks, foundTasks)
	taskRepo.AssertExpectations(t)
}

func TestTaskInteractor_FindByID(t *testing.T) {
	taskRepo := new(MockTaskRepository)
	taskInteractor := NewTaskInteractor(taskRepo)

	task := &entity.Task{ID: 1, Title: "Test Task"}
	taskRepo.On("FindByID", 1).Return(task, nil)

	foundTask, err := taskInteractor.FindByID(1)

	assert.NoError(t, err)
	assert.Equal(t, task, foundTask)
	taskRepo.AssertExpectations(t)
}

func TestTaskInteractor_Update(t *testing.T) {
	taskRepo := new(MockTaskRepository)
	taskInteractor := NewTaskInteractor(taskRepo)

	task := &entity.Task{ID: 1, Title: "Updated Task"}
	taskRepo.On("Update", task).Return(task, nil)

	updatedTask, err := taskInteractor.Update(task)

	assert.NoError(t, err)
	assert.Equal(t, task, updatedTask)
	taskRepo.AssertExpectations(t)
}

func TestTaskInteractor_Delete(t *testing.T) {
	taskRepo := new(MockTaskRepository)
	taskInteractor := NewTaskInteractor(taskRepo)

	task := &entity.Task{ID: 1}
	taskRepo.On("Delete", task).Return(nil)

	err := taskInteractor.Delete(task)

	assert.NoError(t, err)
	taskRepo.AssertExpectations(t)
}
