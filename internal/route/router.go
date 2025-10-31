package route

import (
	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/tasks", GetTasks)
		r.Get("/task/{id}", GetTask)
		r.Post("/tasks", CreateTask)
		r.Put("/task/{id}", UpdateTask)
		r.Delete("/task/{id}", DeleteTask)
	})
}
