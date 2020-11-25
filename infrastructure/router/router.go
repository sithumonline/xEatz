package router

import (
	"github.com/ebonynon/xEatz/usecase/customers"
	"github.com/ebonynon/xEatz/usecase/fooditems"
	"github.com/ebonynon/xEatz/usecase/orders"
	"github.com/ebonynon/xEatz/usecase/restaurants"
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

	r.Post("/orders", orders.AddAOrder)
	r.Get("/orders", orders.GetOrders)
	r.Get("/orders/{id}", orders.GetAOrder)
	r.Delete("/orders/{id}", orders.DeleteAOrder)

	return r
}

// CustomersRouter router for /customers/
func CustomersRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", customers.AddACustomer)
	r.Get("/", customers.GetCustomers)
	r.Get("/{id}", customers.GetACustomer)
	r.Delete("/{id}", customers.DeleteACustomer)

	return r
}

// FoodIteamRouter router for /food-items/
func FoodIteamRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", fooditems.AddAFoodItem)
	r.Get("/", fooditems.GetFoodItems)
	r.Get("/{id}", fooditems.GetAFoodItem)
	r.Delete("/{id}", fooditems.DeleteAFoodItem)

	return r
}

// RestaurantsRouter router for /restaurants/
func RestaurantsRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", restaurants.AddARestaurant)
	r.Get("/", restaurants.GetRestaurants)
	r.Get("/{id}", restaurants.GetARestaurant)
	r.Delete("/{id}", restaurants.DeleteARestaurant)

	return r
}
