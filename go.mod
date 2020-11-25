module github.com/ebonynon/xEatz

go 1.15

replace github.com/ebonynon/xEatz/api/router => ./api/router

replace github.com/ebonynon/xEatz/api/handler => ./api/handler

replace github.com/ebonynon/xEatz/infrastructure/db => ./infrastructure/db

require (
	github.com/go-chi/chi v1.5.0
	github.com/jinzhu/gorm v1.9.16
)
