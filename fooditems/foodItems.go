package fooditems

import "github.com/jinzhu/gorm"

// Restaurants defines the properties of a restaurants
type FoodItems struct {
	gorm.Model
	ItmID      int
	ResID      int
	Name       string
	ItmKind    string
	smallPrice int
	largePrice int
	ImgURL     string
}
