package customers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"

	"github.com/ebonynon/xEatz/db"
)

// OpenDB declare the database
var OpenDB *gorm.DB = db.Database()

// AddACustomer adds a new customer
// POST /addcustomer
func AddACustomer(w http.ResponseWriter, r *http.Request) {

	var customer Customers

	json.NewDecoder(r.Body).Decode(&customer)

	fmt.Println(customer)

	OpenDB.AutoMigrate(&Customers{})
	OpenDB.Create(&customer)

}

// GetCustomers returns customers
// GET /getcustomers
func GetCustomers(w http.ResponseWriter, r *http.Request) {

	var customers []Customers

	OpenDB.Find(&customers)

	json.NewEncoder(w).Encode(&customers)

}

// GetACustomer returns customer
// GET /getcustomer/{id}
func GetACustomer(w http.ResponseWriter, r *http.Request) {

	var customer Customers
	cusID := chi.URLParam(r, "id")

	OpenDB.First(&customer, cusID)

	json.NewEncoder(w).Encode(&customer)

}

// DeleteACustomer drop customer
// DELETE /dropcustomer/{id}
func DeleteACustomer(w http.ResponseWriter, r *http.Request) {

	var customer Customers
	var customers []Customers
	cusID := chi.URLParam(r, "id")

	OpenDB.First(&customer, cusID)
	OpenDB.Delete(&customer)
	OpenDB.Find(&customers)

	json.NewEncoder(w).Encode(&customers)

}
