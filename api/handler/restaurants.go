package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

// Restaurants defines the properties of a restaurants
type Restaurants struct {
	gorm.Model
	ResID    string
	Name     string
	Phone    string
	Location string
	ImgURL   string
}

// AddARestaurant adds a new restaurant
// POST /addrestaurant
func AddARestaurant(w http.ResponseWriter, r *http.Request) {

	var restaurant Restaurants

	json.NewDecoder(r.Body).Decode(&restaurant)

	fmt.Println(restaurant)

	OpenDB.AutoMigrate(&Restaurants{})
	OpenDB.Create(&restaurant)

}

// GetRestaurants returns restaurants
// GET /getrestaurants
func GetRestaurants(w http.ResponseWriter, r *http.Request) {

	var restaurants []Restaurants

	OpenDB.Find(&restaurants)

	json.NewEncoder(w).Encode(&restaurants)

}

// GetARestaurant returns restaurant
// GET /getrestaurant/{id}
func GetARestaurant(w http.ResponseWriter, r *http.Request) {

	var restaurant Restaurants
	restaurantID := chi.URLParam(r, "id")

	OpenDB.First(&restaurant, restaurantID)

	json.NewEncoder(w).Encode(&restaurant)

}

// DeleteARestaurant drop restaurant
// DELETE /droprestaurant/{id}
func DeleteARestaurant(w http.ResponseWriter, r *http.Request) {

	var restaurant Restaurants
	var restaurants []Restaurants
	restaurantID := chi.URLParam(r, "id")

	OpenDB.First(&restaurant, restaurantID)
	OpenDB.Delete(&restaurant)
	OpenDB.Find(&restaurants)

	json.NewEncoder(w).Encode(&restaurants)

}
