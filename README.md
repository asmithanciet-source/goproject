# Task Management REST API Using Go

## Project Overview

The Task Management REST API is a web service developed using the Go programming language. It allows users to create, retrieve, update, and delete tasks through HTTP requests.

The application uses Go's `net/http` package to handle requests and JSON to exchange data. Task information is stored in memory.

## Technologies Used

* Programming Language: Go
* HTTP Server: Go `net/http`
* Data Format: JSON
* Storage: In-memory slice
* Concurrency Control: `sync.Mutex`

## Project Structure

```text
TaskManagementAPI/
├── main.go
├── go.mod
├── README.md
├── models/task.go
├── handlers/task_handler.go
├── middleware/logging.go
└── routes/task_routes.go
```

## Task Fields

| Field       | Type   | Description                                  |
| ----------- | ------ | -------------------------------------------- |
| id          | int    | Identifier assigned by the service           |
| title       | string | Required task name                           |
| description | string | Additional details                           |
| status      | string | Progress status; defaults to `Pending`       |

## API Endpoints

| Method | Endpoint     | Operation        | Result                         |
| ------ | ------------ | ---------------- | ------------------------------ |
| POST   | /tasks       | Create a task    | 201 Created + task JSON        |
| GET    | /tasks       | List all tasks   | 200 OK + JSON array            |
| GET    | /tasks/{id}  | Retrieve a task  | 200 OK; 404 if absent          |
| PUT    | /tasks/{id}  | Update a task    | 200 OK; 404 if absent          |
| DELETE | /tasks/{id}  | Delete a task    | 204 No Content; 404 if absent  |

Other responses: 400 for invalid JSON, missing title or non-numeric ID; 405 for unsupported methods.

## How to Run

```
go run .
```

The server starts at `http://localhost:8080`.

## Testing (PowerShell)

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/tasks" -Method Post -ContentType "application/json" -Body '{"title":"Learn Go REST API","description":"Practice CRUD","status":"Pending"}'
Invoke-RestMethod -Uri "http://localhost:8080/tasks" -Method Get
Invoke-RestMethod -Uri "http://localhost:8080/tasks/1" -Method Get
Invoke-RestMethod -Uri "http://localhost:8080/tasks/1" -Method Put -ContentType "application/json" -Body '{"title":"Complete Go Assignment","description":"Done","status":"Completed"}'
Invoke-WebRequest -Uri "http://localhost:8080/tasks/1" -Method Delete -UseBasicParsing
```

## Validation

```
go build .
go vet ./...
```

## Limitations

* Data is stored only in memory and is lost when the server stops.
* No authentication or authorization.
* No automated tests yet.

## Conclusion

This project demonstrates a modular RESTful web service in Go with CRUD operations, JSON handling, logging middleware, input validation, and mutex-protected shared data.
