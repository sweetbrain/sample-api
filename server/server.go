package main

import "github.com/go-chi/chi"

type Server struct {
	Address string
	Router *chi.Router
}
