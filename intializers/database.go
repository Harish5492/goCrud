package intializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
  )
  var DB *gorm.DB
  func main() {

	var err Error;
	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err !=nil {
		log.Fatal("Error In Connecting the Database")
	}
  }