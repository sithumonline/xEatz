module github.com/ebonynon/xEatz

go 1.15

replace github.com/ebonynon/xEatz/restaurants => ./restaurants

replace github.com/ebonynon/xEatz/db => ./db

replace github.com/ebonynon/xEatz/fooditems => ./fooditems

replace github.com/ebonynon/xEatz/customers => ./customers

replace github.com/ebonynon/xEatz/orders => ./orders

require (
	github.com/go-chi/chi v1.5.0
	github.com/jinzhu/gorm v1.9.16
)
