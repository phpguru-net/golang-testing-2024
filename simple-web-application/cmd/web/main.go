package main

import (
	"fmt"
	"net/http"
	"simplewebapplication/cmd/web/task"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const PORT = 3000

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	apiRoute := chi.NewRouter()
	task.RegisterRoutes(apiRoute)

	r.Mount("/api", apiRoute)

	fmt.Printf("The application is running at http://localhost:%v\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), r)
	return
}
