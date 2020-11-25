package router

import (
	"github.com/ebonynon/xEatz/usecase/handler"
	"github.com/go-chi/chi"
)

// NewRouter A completely separate router for v1 routes
func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Mount("/orders", RestaurantsRouter())
	r.Mount("/customers", RestaurantsRouter())
	r.Mount("/food-items", FoodIteamRouter())
	r.Mount("/restaurants", RestaurantsRouter())

	return r
}

// OrdersRouter router for /orders/
func OrdersRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.AddAOrder)
	r.Get("/", handler.GetOrders)
	r.Get("/{id}", handler.GetAOrder)
	r.Delete("/{id}", handler.DeleteAOrder)

	return r
}

// CustomersRouter router for /customers/
func CustomersRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.AddACustomer)
	r.Get("/", handler.GetCustomers)
	r.Get("/{id}", handler.GetACustomer)
	r.Delete("/{id}", handler.DeleteACustomer)

	return r
}

// FoodIteamRouter router for /food-items/
func FoodIteamRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.AddAFoodItem)
	r.Get("/", handler.GetFoodItems)
	r.Get("/{id}", handler.GetAFoodItem)
	r.Delete("/{id}", handler.DeleteAFoodItem)

	return r
}

// RestaurantsRouter router for /restaurants/
func RestaurantsRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.AddARestaurant)
	r.Get("/", handler.GetRestaurants)
	r.Get("/{id}", handler.GetARestaurant)
	r.Delete("/{id}", handler.DeleteARestaurant)

	return r
}
