package handler

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

type Method string

const (
	POST   Method = "POST"
	GET    Method = "GET"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
	PATCH  Method = "PATCH"
)

// Route is the information for every URI.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc fiber.Handler
}


type Routes []Route

func AddServices(engine *fiber.App) fiber.Router {
	fmt.Printf("Entering TaskRoutesServices")
	taskgroup := engine.Group("/tasks")
	// Group for Application Data API
	for _, route := range taskRoutes {
		registerRoute(taskgroup, route)
	}

	fmt.Print("\n AddService exit")

	return taskgroup
}

// Helper function to register routes
func registerRoute(group fiber.Router, route Route) {
	switch route.Method {
	case "GET":
		group.Get(route.Pattern, route.HandlerFunc)
	case "POST":
		group.Post(route.Pattern, route.HandlerFunc)
	case "PUT":
		group.Put(route.Pattern, route.HandlerFunc)
	case "PATCH":
		group.Patch(route.Pattern, route.HandlerFunc)
	case "DELETE":
		group.Delete(route.Pattern, route.HandlerFunc)
	}
}

var taskRoutes = Routes{

	// Authentication Data
	{
		"GetTaskList",
		http.MethodGet,
		"/alltasks",
		GetTaskListHandler,
	},

	// Authentication Subscription
	{
		"PatchAuthenticationSubscription",
		http.MethodPost,
		"/creatTask",
		CreatTaskHandler,
	},
}