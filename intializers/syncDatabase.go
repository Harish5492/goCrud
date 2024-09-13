package intializers

import "github.com/HarishRana/goCrud/src/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})

	// add an initial user
	AddInitialUser()
}

func AddInitialUser() {
	DB.Create(&models.User{FullName:"Harish Rana", Email: "harishrana5492@gmail.com", Password: "12345678"})
}
