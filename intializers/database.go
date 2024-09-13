package intializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"log"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := os.Getenv("DATABASE")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
}
