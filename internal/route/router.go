package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/hive-bootcamp/go-rest-api-homework/internal/handlers"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", handlers.GetTasks)
		r.Get("/{id}", handlers.GetTask)
		r.Post("/", handlers.CreateTask)
		r.Put("/{id}", handlers.UpdateTask)
		r.Delete("/{id}", handlers.DeleteTask)
	})

	return r
}
