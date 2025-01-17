package router

import (
	"dattis/controllers/customers"
	"dattis/controllers/invoices"
	"dattis/controllers/tax"
	"github.com/go-chi/chi/v5"
)

func GetRoutes(r *chi.Mux) *chi.Mux {
	r.Route("/api", func(r chi.Router) {
		r.Get("/invoices", invoices.Index)
		r.Post("/invoices", invoices.Create)
		r.Put("/invoices/{id}", invoices.Update)
		r.Get("/invoices/{id}", invoices.Show)

		r.Get("/customers", customers.Index)

		r.Get("/tax-rates", tax.Index)
	})

	return r
}
