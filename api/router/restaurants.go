package router

import (
	"github.com/ebonynon/xEatz/api/handler"
	"github.com/go-chi/chi"
)

// RestaurantsRouter router for /restaurants/
func RestaurantsRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.AddARestaurant)
	r.Get("/", handler.GetRestaurants)
	r.Get("/{id}", handler.GetARestaurant)
	r.Put("/{id}", handler.UpdateARestaurant)
	r.Delete("/{id}", handler.DeleteARestaurant)

	return r
}
