package handler

import "github.com/jinzhu/gorm"

// Restaurants defines the properties of a restaurants
type Restaurants struct {
	gorm.Model
	ResID    int
	Name     string
	Phone    int
	Location string
	ImgURL   string
}
