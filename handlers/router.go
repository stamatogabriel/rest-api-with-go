package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Response struct {
	Msg  string
	Code int
}

func CreateRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/health", healthCheckHandler)

			/**	TODO Handlers */
			r.Get("/todos", getTodosHandler)
			r.Get("/todos/{id}", getTodoByIDHandler)
			r.Post("/todos/create", createTodoHandler)
			r.Put("/todos/update/{id}", updateTodoHandler)
			r.Delete("/todos/delete/{id}", deleteTodoHandler)
		})
	})

	return router
}
