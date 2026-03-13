package handlers


import (
	"log"
	"github.com/fc637/go-task-api/internal/models"
	"github.com/fc637/go-task-api/internal/taskstore"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TaskHandler struct {
	store *store.TaskStore
}

func NewTaskHandler(store *store.TaskStore) *TaskHandler {
	return &TaskHandler{store: store}
}


func (h *TaskHandler) CreateTaskHandler(c *fiber.Ctx) error {
	log.Print("Entering Into Create Task Handler \n")
	defer log.Print("\nExiting Create Task Handler\n")

	// req := new(models.CreateTaskRequest)
	var req models.CreateTaskRequest

	if err := c.BodyParser(&req); err != nil {
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
	}else{
		req.Status= string(models.Todo)
		status=models.TaskStatus(req.Status)
	}

	task := models.Task{
		ID:     uuid.New().String(),
		Title:  req.Title,
		Status: status,
	}

	h.store.Add(task)

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandler) GetTaskListHandler(c *fiber.Ctx) error {
	log.Print("Entering Into Get Handler List\n")
	log.Print("Exiting Get Handler List\n")
	tasks := h.store.List()

	return c.JSON(tasks)
}