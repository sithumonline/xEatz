package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/ebonynon/xEatz/customers"
	"github.com/ebonynon/xEatz/fooditems"
	"github.com/ebonynon/xEatz/orders"
	"github.com/ebonynon/xEatz/restaurants"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	getPort := func(KEY string) string {

		ENV, Okay := os.LookupEnv(KEY)

		if !Okay {

			ENV = "3080"

		}
		return ENV

	}

	r.Post	("/orders", orders.AddAOrder)
	r.Get	("/orders", orders.GetOrders)
	r.Get	("/orders/{id}", orders.GetAOrder)
	r.Delete("/orders/{id}", orders.DeleteAOrder)

	r.Post	("/customers", customers.AddACustomer)
	r.Get	("/customers", customers.GetCustomers)
	r.Get	("/customers/{id}", customers.GetACustomer)
	r.Delete("/customer/{id}", customers.DeleteACustomer)

	r.Post	("/food-items", fooditems.AddAFoodItem)
	r.Get	("/food-items", fooditems.GetFoodItems)
	r.Get	("/food-items/{id}", fooditems.GetAFoodItem)
	r.Delete("/food-items/{id}", fooditems.DeleteAFoodItem)

	r.Post	("/restaurants", restaurants.AddARestaurant)
	r.Get	("/restaurants", restaurants.GetRestaurants)
	r.Get	("/restaurants/{id}", restaurants.GetARestaurant)
	r.Delete("/restaurants/{id}", restaurants.DeleteARestaurant)
	
	http.ListenAndServe(":"+getPort("PORT"), r)

}
