module github.com/ebonynon/xEatz

go 1.15

replace github.com/ebonynon/xEatz/api/router => ./api/router

replace github.com/ebonynon/xEatz/api/handler => ./api/handler

replace github.com/ebonynon/xEatz/infrastructure/db => ./infrastructure/db

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/go-chi/chi v1.5.0
	github.com/go-chi/cors v1.1.1
	github.com/jinzhu/gorm v1.9.16
)
