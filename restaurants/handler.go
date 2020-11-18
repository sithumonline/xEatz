package restaurants

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ebonynon/xEatz/db"
)

// AddRestaurants adds a new restaurants
// POST /addrestaurants
func AddRestaurants(w http.ResponseWriter, r *http.Request) {

	var restaurants Restaurants
	OpenDB := db.Database()

	json.NewDecoder(r.Body).Decode(&restaurants)

	fmt.Println(restaurants)

	OpenDB.AutoMigrate(&Restaurants{})
	OpenDB.Create(&restaurants)

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
