package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

var err error

func Database() *gorm.DB {

	getURL := func(KEY string) string {

		ENV, Okay := os.LookupEnv(KEY)
		URL := "host=0.0.0.0 port=5432 user=postgres dbname=postgres sslmode=disable password=password"

		if !Okay {

			ENV = URL

		}
		return ENV

	}

	DB, err = gorm.Open("postgres", getURL("DATABASE_URL"))

	if err != nil {

		fmt.Println(err)
		panic("failed to connect database")

	} else {

		fmt.Println("database connected")

	}

	return DB
}
