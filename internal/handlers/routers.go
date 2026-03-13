package handlers

import (

	"github.com/gofiber/fiber/v2"
	"github.com/fc637/go-task-api/internal/taskstore"
	"log"
)


func AddServices(engine *fiber.App) fiber.Router {

	log.Println("Entering TaskRoutesServices")

	taskgroup := engine.Group("/tasks")

	store := store.NewTaskStore()
	handler := NewTaskHandler(store)

	taskgroup.Get("/alltasks", handler.GetTaskListHandler)
	taskgroup.Post("/creatTask", handler.CreateTaskHandler)

	log.Println("AddService exit")

	return taskgroup
}
