package router

import (
	"github.com/ebonynon/xEatz/api/handler"
	"github.com/go-chi/chi"
)

// OrdersRouter router for /orders/
func OrdersRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.AddAOrder)
	r.Get("/", handler.GetOrders)
	r.Get("/{id}", handler.GetAOrder)
	r.Delete("/{id}", handler.DeleteAOrder)

	return r
}
