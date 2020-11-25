package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/ebonynon/xEatz/api/router"
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

	r.Mount("/v1", router.NewRouter())

	http.ListenAndServe(":"+getPort("PORT"), r)

}
