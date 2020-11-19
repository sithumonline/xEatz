package orders

import "github.com/jinzhu/gorm"

// Orders defines the properties of a orders
type Orders struct {
	gorm.Model
	CusID      int
	OrdID      int
	ResID      int
	NoItems    int
	paymentWay string
}
