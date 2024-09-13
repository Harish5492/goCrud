package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/HarishRana/goCrud/intializers"
	"github.com/HarishRana/goCrud/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// get email and password from request body
	var body struct {
		FullName string
		Email    string
		Password string
	}

	if c.Bind(&body) == nil {
		models.ReturnGenericBadResponse(c, "Invalid request body 1")
		return
	}

	if (body.FullName == "" || body.Email == "") || (body.Password == "") {
		models.ReturnGenericBadResponse(c, " FullName, Email and Password cannot be empty")
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		models.ReturnGenericBadResponse(c, "Failed to hash password")
		return
	}
	// create user
	user := models.User{FullName: body.FullName, Email: body.Email, Password: string(hash)}
	// add to table
	result := initalizers.DB.Create(&user)

	if result.Error != nil {
		models.ReturnGenericBadResponse(c, "Failed to create user")
		return
	}

	// return response
	models.ReturnGenericSuccessWithNoMessageResponse(c)
}