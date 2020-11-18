package restaurants

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ebonynon/xEatz/db"
)

// AddARestaurant adds a new restaurant
// POST /addrestaurant
func AddARestaurant(w http.ResponseWriter, r *http.Request) {

	var restaurant Restaurants
	OpenDB := db.Database()

	json.NewDecoder(r.Body).Decode(&restaurant)

	fmt.Println(restaurant)

	OpenDB.AutoMigrate(&Restaurants{})
	OpenDB.Create(&restaurant)

	defer OpenDB.Close()
}

// GetRestaurants returns restaurants
// GET /getrestaurants
func GetRestaurants(w http.ResponseWriter, r *http.Request) {

	var restaurants []Restaurants
	OpenDB := db.Database()

	OpenDB.Find(&restaurants)

	json.NewEncoder(w).Encode(&restaurants)

}

// GetARestaurant returns restaurant
// GET /getrestaurant
func GetARestaurant(w http.ResponseWriter, r *http.Request) {

	var restaurant Restaurants
	restaurantID := chi.URLParam(r, "id")

	OpenDB := db.Database()

	OpenDB.First(&restaurant, restaurantID)

	json.NewEncoder(w).Encode(&restaurant)

}
