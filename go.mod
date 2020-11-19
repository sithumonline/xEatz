module github.com/ebonynon/xEatz

go 1.15

replace github.com/ebonynon/xEatz/restaurants => ./restaurants

replace github.com/ebonynon/xEatz/db => ./db

replace github.com/ebonynon/xEatz/fooditems => ./fooditems

replace github.com/ebonynon/xEatz/customers => ./customers

replace github.com/ebonynon/xEatz/orders => ./orders

require (
	github.com/ebonynon/xEatz/customers v0.0.0-00010101000000-000000000000
	github.com/ebonynon/xEatz/db v0.0.0-00010101000000-000000000000 // indirect
	github.com/ebonynon/xEatz/fooditems v0.0.0-00010101000000-000000000000
	github.com/ebonynon/xEatz/orders v0.0.0-00010101000000-000000000000
	github.com/ebonynon/xEatz/restaurants v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi v1.5.0
	github.com/jinzhu/gorm v1.9.16
)
