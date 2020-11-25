package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// OpenDB declare the database
//var OpenDB *gorm.DB = db.Database()

// AddAFoodItem adds a new food Item
// POST /addfooditem
func AddAFoodItem(w http.ResponseWriter, r *http.Request) {

	var foodItem FoodItems

	json.NewDecoder(r.Body).Decode(&foodItem)

	fmt.Println(foodItem)

	OpenDB.AutoMigrate(&FoodItems{})
	OpenDB.Create(&foodItem)

}

// GetFoodItems returns food Item
// GET /getfooditems
func GetFoodItems(w http.ResponseWriter, r *http.Request) {

	var foodItems []FoodItems

	OpenDB.Find(&foodItems)

	json.NewEncoder(w).Encode(&foodItems)

}

// GetAFoodItem returns food Item
// GET /getfooditem/{id}
func GetAFoodItem(w http.ResponseWriter, r *http.Request) {

	var foodItem FoodItems
	foodID := chi.URLParam(r, "id")

	OpenDB.First(&foodItem, foodID)

	json.NewEncoder(w).Encode(&foodItem)

}

// DeleteAFoodItem drop restaurant
// DELETE /dropfooditem/{id}
func DeleteAFoodItem(w http.ResponseWriter, r *http.Request) {

	var foodItem FoodItems
	var foodItems []FoodItems
	foodID := chi.URLParam(r, "id")

	OpenDB.First(&foodItem, foodID)
	OpenDB.Delete(&foodItem)
	OpenDB.Find(&foodItems)

	json.NewEncoder(w).Encode(&foodItems)

}
