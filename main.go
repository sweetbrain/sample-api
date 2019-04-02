package main

import (
	"github.com/go-chi/chi"
	"net/http"

	"./handler"
)

func main() {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.UserRegister)
		r.Get("/", handler.UserReader)
		r.Put("/", handler.UserUpdater)
		r.Delete("/", handler.UserDeleter)
	})

	http.ListenAndServe(":8080", r)
}
