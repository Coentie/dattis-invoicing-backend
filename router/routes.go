package router

import (
	"dattis/controllers/invoices"
	"github.com/go-chi/chi/v5"
)

func GetRoutes(r *chi.Mux) *chi.Mux {
	r.Get("/invoices", invoices.Index)
	r.Put("/invoices/{id}", invoices.Update)

	return r
}