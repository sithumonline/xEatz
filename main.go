package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

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

	r.Post("/addrestaurant", restaurants.AddARestaurant)
	r.Get("/getrestaurants", restaurants.GetRestaurants)
	r.Get("/getrestaurant/{id}", restaurants.GetARestaurant)
	r.Delete("/droprestaurant/{id}", restaurants.DeleteARestaurant)
	http.ListenAndServe(":"+getPort("PORT"), r)
}
