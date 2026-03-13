# GO Task API

A simple REST API built using **Golang** and the **Fiber framework** to manage tasks.

This project allows users to:

* Create tasks
* Retrieve task list

The tasks are stored in an **in-memory store**.


# Tech Stack

* Golang
* Fiber Framework
* UUID for task IDs
* In-memory task storage



# Project Structure

```
go-task-api
│
├── cmd
│   └── main.go            # Application entry point
│
├── internal
│   ├── handlers           # API handlers and routes
│   ├── models             # Task models
│   └── taskstore          # In-memory data store
│
├── tests                  # Unit tests
│
├── go.mod
└── README.md
```

# Install Dependencies

Run the following command:

go mod tidy

# Run the Application

From the project root:   go run ./cmd/main.go


Server will start at:  http://localhost:8000

# API Endpoints

## 1. Create Task

### Endpoint

```
POST /tasks/creatTask
```

### Curl Request


curl -X POST http://localhost:8000/tasks/creatTask -H "Content-Type: application/json" 
-d '{
"title":"Learn Golang",
"status":"todo"
}'


### Request Body

| Field  | Type   | Required | Description                     |
| ------ | ------ | -------- | ------------------------------- |
| title  | string | Yes      | Task title (max 200 characters) |
| status | string | No       | todo / in_progress / done       |

If status is omitted, default is **todo**.

---

### Example Response

```
{
"id": "b8b5b77c-22d5-4f51-9e45-4c2d9a01c6f3",
"title": "Learn Golang",
"status": "todo"
}
```

---

## 2. Get All Tasks

### Endpoint

```
GET /tasks/alltasks
```

### Curl Request

```
curl http://localhost:8000/tasks/alltasks
```

### Example Response

```
[
{
"id": "b8b5b77c-22d5-4f51-9e45-4c2d9a01c6f3",
"title": "Learn Golang",
"status": "todo"
},
{
"id": "c1b9b4b7-3c67-49e8-9e45-1ab4e0b0a92a",
"title": "Build REST API",
"status": "in_progress"
}
]
```

---

# Validation Rules

* Title is **required**
* Title length must be **≤ 200 characters**
* Status must be one of:

  * `todo`
  * `in_progress`
  * `done`

---

# Error Example

### Missing Title

```
curl -X POST http://localhost:8000/tasks/creatTask \
-H "Content-Type: application/json" \
-d '{}'
```

Response

```
{
"error": "title is required"
}
```

---

# Run Tests

go test ./tests -v


---

# Example Output
=== RUN   TestCreateTask
2026/03/14 02:25:00 Entering Into Create Task Handler 
2026/03/14 02:25:00 
Exiting Create Task Handler
--- PASS: TestCreateTask (0.00s)
=== RUN   TestCreateTaskTitleTooLong
2026/03/14 02:25:00 Entering Into Create Task Handler 
2026/03/14 02:25:00 
Exiting Create Task Handler
--- PASS: TestCreateTaskTitleTooLong (0.00s)
=== RUN   TestCreateTaskMissingTitle
2026/03/14 02:25:00 Entering Into Create Task Handler 
2026/03/14 02:25:00 
Exiting Create Task Handler
--- PASS: TestCreateTaskMissingTitle (0.00s)
=== RUN   TestCreateTaskDefaultStatus
2026/03/14 02:25:00 Entering Into Create Task Handler 
2026/03/14 02:25:00 
Exiting Create Task Handler
--- PASS: TestCreateTaskDefaultStatus (0.00s)
=== RUN   TestCreateTaskInvalidStatus
2026/03/14 02:25:00 Entering Into Create Task Handler 
2026/03/14 02:25:00 
Exiting Create Task Handler
--- PASS: TestCreateTaskInvalidStatus (0.00s)
=== RUN   TestGetTasks
2026/03/14 02:25:00 Entering Into Get Handler List
2026/03/14 02:25:00 Exiting Get Handler List
--- PASS: TestGetTasks (0.00s)
PASS
ok  	command-line-arguments	0.005s

