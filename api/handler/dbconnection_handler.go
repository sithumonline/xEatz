package handler

import (
	"github.com/ebonynon/xEatz/infrastructure/db"
	"github.com/jinzhu/gorm"
)

// OpenDB declare the database
var OpenDB *gorm.DB = db.Database()
