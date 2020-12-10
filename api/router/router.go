package router

import (
	"github.com/go-chi/chi"
)

// NewRouter A completely separate router for v1 routes
func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Mount("/orders", OrdersRouter())
	r.Mount("/customers", CustomersRouter())
	r.Mount("/food-items", FoodIteamRouter())
	r.Mount("/restaurants", RestaurantsRouter())

	return r
}
