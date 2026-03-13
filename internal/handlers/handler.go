package handler


import (
	"github.com/fc637/go-task-api/internal/models"
	"github.com/fc637/go-task-api/internal/store"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TaskHandler struct {
	store *store.TaskStore
}

func NewTaskHandler(store *store.TaskStore) *TaskHandler {
	return &TaskHandler{store: store}
}


func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	fmt.Printf("Entering Into Create Task Handler")
	defer fmt.Printf("Exiting Create Task Handler")

	// req := new(models.CreateTaskRequest)
	var req models.CreateTaskRequest

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body :" +err.Error()})
	}

	if req.Title == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "title is required"})
	}

	if len(req.Title) > 200 {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "title must be <= 200 characters"})
	}

	// status := models.Todo
	var status models.TaskStatus
	if req.Status != "" {
		switch req.Status {
		case "todo", "in_progress", "done":
			status = models.TaskStatus(req.Status)
		default:
			return c.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"error": "invalid status"})
		}
	}

	task := models.Task{
		ID:     uuid.New().String(),
		Title:  req.Title,
		Status: status,
	}

	h.store.Add(task)

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandler) ListTasks(c *fiber.Ctx) error {
	fmt.Printf("Entering Into Get Handler List")
	defer fmt.Printf("Exiting Get Handler List")
	tasks := h.store.List()

	return c.JSON(tasks)
}