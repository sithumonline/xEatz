package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

// Orders defines the properties of a orders
type Orders struct {
	gorm.Model
	CusID      int
	OrdID      int
	ResID      int
	NoItems    int
	paymentWay string
}

// AddAOrder adds a new order
// POST /addorder
func AddAOrder(w http.ResponseWriter, r *http.Request) {

	var order Orders

	json.NewDecoder(r.Body).Decode(&order)

	fmt.Println(order)

	OpenDB.AutoMigrate(&Orders{})
	OpenDB.Create(&order)

}

// GetOrders returns orders
// GET /getorders
func GetOrders(w http.ResponseWriter, r *http.Request) {

	var orders []Orders

	OpenDB.Find(&orders)

	json.NewEncoder(w).Encode(&orders)

}

// GetAOrder returns order
// GET /getorder/{id}
func GetAOrder(w http.ResponseWriter, r *http.Request) {

	var order Orders
	ordID := chi.URLParam(r, "id")

	OpenDB.First(&order, ordID)

	json.NewEncoder(w).Encode(&order)

}

// DeleteAOrder drop order
// DELETE /droporder/{id}
func DeleteAOrder(w http.ResponseWriter, r *http.Request) {

	var order Orders
	var orders []Orders
	ordID := chi.URLParam(r, "id")

	OpenDB.First(&order, ordID)
	OpenDB.Delete(&order)
	OpenDB.Find(&orders)

	json.NewEncoder(w).Encode(&orders)

}

// UpdateAOrder update order
// PUT /updateorder/{id}
func UpdateAOrder(w http.ResponseWriter, r *http.Request) {

	var order Orders
	var orders []Orders
	ordID := chi.URLParam(r, "id")

	json.NewDecoder(r.Body).Decode(&order)

	fmt.Println(order)

	OpenDB.Model(&order).Where("ID = ?", ordID).Updates(order)
	OpenDB.Find(&orders)

	json.NewEncoder(w).Encode(&orders)

}
