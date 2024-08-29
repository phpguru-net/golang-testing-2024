package task

import (
	"fmt"
	"net/http"
	"strconv"

	"simplewebapplication/cmd/web/core"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	// r.Get("/tasks", getTasks)
	// r.Get("/tasks/{id}", getTask)
	r.Route("tasks", func(r chi.Router) {
		r.Get("/", getTasks)
		r.Get("/tasks/{id}", getTask)
	})
}

type Task struct {
	Title string `json:title`
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	// Append new tasks to the slice
	tasks = append(tasks, Task{
		Title: "Task 1",
	})
	tasks = append(tasks, Task{
		Title: "Task 2",
	})

	core.HttpResponseSuccess(w, tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		core.HttpResponseBadRequest(w, err)
		return
	}

	var task Task = Task{
		Title: fmt.Sprintf("Task %v", id),
	}

	core.HttpResponseSuccess(w, task)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		core.HttpResponseBadRequest(w, err)
		return
	}

	var task Task = Task{
		Title: fmt.Sprintf("Task %v", id),
	}

	core.HttpResponseCreated(w, task)
}
