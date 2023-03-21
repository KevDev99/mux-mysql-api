package configs

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", EnvMySQL())

	if err != nil {
		log.Fatal(err)
	}

	return db

}

// Client instance
var DB *gorm.DB = ConnectDB()
