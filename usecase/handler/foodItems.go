package handler

import "github.com/jinzhu/gorm"

// FoodItems defines the properties of a fooditems
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
