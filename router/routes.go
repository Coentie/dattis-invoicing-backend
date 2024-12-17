package router

import (
	"dattis/controllers/customers"
	"dattis/controllers/invoices"
	"github.com/go-chi/chi/v5"
)

func GetRoutes(r *chi.Mux) *chi.Mux {
	r.Route("/api", func(r chi.Router) {
		r.Get("/invoices", invoices.Index)
		r.Put("/invoices/{id}", invoices.Update)

		r.Get("/customers", customers.Index)
	})

	return r
}
