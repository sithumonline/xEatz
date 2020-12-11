package router

import (
	"github.com/ebonynon/xEatz/api/handler"
	"github.com/go-chi/chi"
)

// FoodIteamRouter router for /food-items/
func FoodIteamRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.AddAFoodItem)
	r.Get("/", handler.GetFoodItems)
	r.Get("/{id}", handler.GetAFoodItem)
	r.Put("/{id}", handler.UpdateAFoodItem)
	r.Delete("/{id}", handler.DeleteAFoodItem)

	return r
}
