package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/ebonynon/xEatz/api/router"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	getPort := func(KEY string) string {

		ENV, Okay := os.LookupEnv(KEY)

		if !Okay {

			ENV = "3080"

		}
		return ENV

	}

	r.Mount("/v1", router.NewRouter())

	http.ListenAndServe(":"+getPort("PORT"), r)

}
