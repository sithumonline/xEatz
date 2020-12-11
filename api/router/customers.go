package router

import (
	"github.com/ebonynon/xEatz/api/handler"
	"github.com/go-chi/chi"
)

// CustomersRouter router for /customers/
func CustomersRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.AddACustomer)
	r.Get("/", handler.GetCustomers)
	r.Get("/{id}", handler.GetACustomer)
	r.Put("/{id}", handler.UpdateACustomer)
	r.Delete("/{id}", handler.DeleteACustomer)

	return r
}
