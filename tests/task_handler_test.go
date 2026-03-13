package tests

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/fc637/go-task-api/internal/handlers"
	"github.com/fc637/go-task-api/internal/taskstore"

	"github.com/gofiber/fiber/v2"
)
func TestCreateTask(t *testing.T) {

	app := fiber.New()

	store := store.NewTaskStore()
	handler := handlers.NewTaskHandler(store)

	app.Post("/tasks", handler.CreateTaskHandler)

	req := httptest.NewRequest(
		"POST",
		"/tasks",
		bytes.NewBuffer([]byte(`{"title":"test task"}`)),
	)

	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusCreated {
		t.Errorf("expected 201 got %d", resp.StatusCode)
	}
}
func TestCreateTaskTitleTooLong(t *testing.T) {

	app := fiber.New()

	store := store.NewTaskStore()
	handler := handlers.NewTaskHandler(store)

	app.Post("/tasks", handler.CreateTaskHandler)

	longTitle := make([]byte, 201)
	for i := range longTitle {
		longTitle[i] = 'a'
	}

	body := `{"title":"` + string(longTitle) + `"}`

	req := httptest.NewRequest(
		"POST",
		"/tasks",
		bytes.NewBuffer([]byte(body)),
	)

	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Errorf("expected 400 got %d", resp.StatusCode)
	}
}

func TestCreateTaskMissingTitle(t *testing.T) {

	app := fiber.New()

	store := store.NewTaskStore()
	handler := handlers.NewTaskHandler(store)

	app.Post("/tasks", handler.CreateTaskHandler)

	req := httptest.NewRequest(
		"POST",
		"/tasks",
		bytes.NewBuffer([]byte(`{}`)),
	)

	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Errorf("expected 400 got %d", resp.StatusCode)
	}
}

func TestCreateTaskDefaultStatus(t *testing.T) {

	app := fiber.New()

	store := store.NewTaskStore()
	handler := handlers.NewTaskHandler(store)

	app.Post("/tasks", handler.CreateTaskHandler)

	req := httptest.NewRequest(
		"POST",
		"/tasks",
		bytes.NewBuffer([]byte(`{"title":"default status test"}`)),
	)

	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusCreated {
		t.Errorf("expected 201 got %d", resp.StatusCode)
	}
}

func TestCreateTaskInvalidStatus(t *testing.T) {

	app := fiber.New()

	store := store.NewTaskStore()
	handler := handlers.NewTaskHandler(store)

	app.Post("/tasks", handler.CreateTaskHandler)

	req := httptest.NewRequest(
		"POST",
		"/tasks",
		bytes.NewBuffer([]byte(`{"title":"task","status":"invalid"}`)),
	)

	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Errorf("expected 400 got %d", resp.StatusCode)
	}
}

func TestGetTasks(t *testing.T) {

	app := fiber.New()

	store := store.NewTaskStore()
	handler := handlers.NewTaskHandler(store)

	app.Get("/tasks", handler.GetTaskListHandler)

	req := httptest.NewRequest(
		"GET",
		"/tasks",
		nil,
	)

	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected 200 got %d", resp.StatusCode)
	}
}



