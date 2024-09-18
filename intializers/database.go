package intializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/HarishRana/goCrud/models" // Import the models package
)

var DB *gorm.DB

func ConnectToDatabase() {
	dsn := os.Getenv("DATABASE")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Run the migration to create the `users` table
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}
}
