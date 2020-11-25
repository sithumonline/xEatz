package handler

import "github.com/jinzhu/gorm"

// Customers defines the properties of a customers
type Customers struct {
	gorm.Model
	CusID int
	FName string
	LName string
	Phone int
	Mail  string
}
