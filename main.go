package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Post("/", UserRegister)
		r.Get("/", UserReader)
		r.Put("/", UserUpdater)
		r.Delete("/", UserDeleter)
	})

	http.ListenAndServe(":8080", r)
}