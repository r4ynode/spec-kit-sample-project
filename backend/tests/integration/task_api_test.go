package integration

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"backend/src/entity"
	"backend/src/router"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&entity.Task{})
	return db
}

func TestCreateTask(t *testing.T) {
	db := setupTestDB()
	r := mux.NewRouter()
	router.SetupRoutes(r, db) // Use the SetupRoutes from the router package

	var jsonStr = []byte(`{"title":"Test Task"}`)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expected := `{"id":1,"title":"Test Task","completed":false` // ID and CreatedAt will vary
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetTasks(t *testing.T) {
	db := setupTestDB()
	r := mux.NewRouter()
	router.SetupRoutes(r, db)

	// Create a task first
	db.Create(&entity.Task{Title: "Task 1", Completed: false})

	req, _ := http.NewRequest("GET", "/tasks", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":1,"title":"Task 1","completed":false`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUpdateTask(t *testing.T) {
	db := setupTestDB()
	r := mux.NewRouter()
	router.SetupRoutes(r, db)

	// Create a task first
	task := entity.Task{Title: "Task to Update", Completed: false}
	db.Create(&task)

	var jsonStr = []byte(`{"title":"Updated Task","completed":true}`)
	req, _ := http.NewRequest("PUT", "/tasks/"+strconv.Itoa(task.ID), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"id":1,"title":"Updated Task","completed":true`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDeleteTask(t *testing.T) {
	db := setupTestDB()
	r := mux.NewRouter()
	router.SetupRoutes(r, db)

	// Create a task first
	task := entity.Task{Title: "Task to Delete", Completed: false}
	db.Create(&task)

	req, _ := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(task.ID), nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}

	// Verify task is deleted
	var deletedTask entity.Task
	if err := db.First(&deletedTask, task.ID).Error; err == nil {
		t.Errorf("Task was not deleted")
	}
}
