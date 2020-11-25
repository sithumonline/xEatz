module github.com/ebonynon/xEatz

go 1.15

replace github.com/ebonynon/xEatz/usecase/orders => ./usecase/orders

replace github.com/ebonynon/xEatz/usecase/customers => ./usecase/customers

replace github.com/ebonynon/xEatz/usecase/fooditems => ./usecase/fooditems

replace github.com/ebonynon/xEatz/usecase/restaurants => ./usecase/restaurants

replace github.com/ebonynon/xEatz/infrastructure/db => ./infrastructure/db

replace github.com/ebonynon/xEatz/infrastructure/router => ./infrastructure/router


require (
	github.com/go-chi/chi v1.5.0
	github.com/jinzhu/gorm v1.9.16
)
