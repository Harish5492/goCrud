package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/HarishRana/goCrud/common"
	"github.com/HarishRana/goCrud/controller"
	"github.com/HarishRana/goCrud/intializers"
)

// init -> run before main
func init() {
	intializers.LoadEnvVariables()
	intializers.ConnectToDatabase()
	// initialize logger
	common.SetupLogger()
}

func main() {
	fmt.Println("Hello Go")
	fmt.Println("It is working")

	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	r.Run() // listen and serve on 0.0.0.0:8080
}
