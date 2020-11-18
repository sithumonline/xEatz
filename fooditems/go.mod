module github.com/ebonynon/xEatz/fooditems

go 1.15

replace github.com/ebonynon/xEatz/db => ../db

require (
	github.com/ebonynon/xEatz/db v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi v1.5.0
	github.com/jinzhu/gorm v1.9.16
)
