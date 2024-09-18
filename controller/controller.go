package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/HarishRana/goCrud/common"
	"github.com/HarishRana/goCrud/intializers"
	"github.com/HarishRana/goCrud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Signup function for user registration
func Signup(c *gin.Context) {
	// Get email and password from request body
	var body struct {
		FullName string `json:"FullName"`
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}

	// Bind the request body and log it
	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println("Request Body:", body)
		common.ReturnGenericBadResponse(c, "Invalid request body")
		return
	}

	// Log request body for debugging
	fmt.Printf("Signup Request Body: %+v\n", body)

	// Validate the input fields
	if (body.FullName == "" || body.Email == "") || (body.Password == "") {
		common.ReturnGenericBadResponse(c, "FullName, Email, and Password cannot be empty")
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		common.ReturnGenericBadResponse(c, "Failed to hash password")
		return
	}

	// Create user
	user := models.User{FullName: body.FullName, Email: body.Email, Password: string(hash)}
	fmt.Printf("usere here is",user)
	fmt.Println("baby doll main sone di",user)

	// Add to the database
	result := intializers.DB.Create(&user)
	if result.Error != nil {
		common.ReturnGenericBadResponse(c, "Failed to create user")
		return
	}

	// Return response
	common.ReturnGenericSuccessWithNoMessageResponse(c)
}

// Login function for user authentication
func Login(c *gin.Context) {
	// Get email and password from request body
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind the request body and log it
	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println("Request Body:", body)
		common.ReturnGenericBadResponse(c, "Invalid request body")
		return
	}

	// Log request body for debugging
	fmt.Printf("Login Request Body: %+v\n", body)

	// Validate input fields
	if (body.Email == "") || (body.Password == "") {
		common.ReturnGenericBadResponse(c, "Email and Password cannot be empty")
		return
	}

	// Look up email in the database
	var user models.User
	intializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		// Didn't find user
		common.ReturnGenericBadResponse(c, "Invalid email or password")
		return
	}

	// Compare the hashed password with the one sent
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		common.ReturnGenericBadResponse(c, "Invalid email or password")
		return
	}

	// Generate a JWT token and send it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign the token using the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		common.ReturnGenericBadResponse(c, "Failed to create token")
		return
	}

	// Set token as a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// Send success response
	common.ReturnGenericSuccessResponse(c,tokenString)
}

func UpdateUser(c *gin.Context) {
	// Get the user ID from URL parameters
	id := c.Param("id")

	// Get the updated data from the request body
	var body struct {
		FullName string
		Email    string
		Password string
	}

	// Bind the request body to the struct
	if err := c.Bind(&body); err != nil {
		common.ReturnGenericBadResponse(c, "Invalid request body")
		return
	}

	// Find the user by ID
	var user models.User
	if result := intializers.DB.First(&user, "id = ?", id); result.Error != nil {
		common.ReturnGenericBadResponse(c, "User not found")
		return
	}

	// Update the fields if provided
	if body.FullName != "" {
		user.FullName = body.FullName
	}
	if body.Email != "" {
		user.Email = body.Email
	}
	if body.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			common.ReturnGenericBadResponse(c, "Failed to hash password")
			return
		}
		user.Password = string(hash)
	}

	// Save the updated user data
	if result := intializers.DB.Save(&user); result.Error != nil {
		common.ReturnGenericBadResponse(c, "Failed to update user")
		return
	}

	// Return success response
	common.ReturnGenericSuccessResponse(c, "Updated Successfully")
}
func DeleteUser(c *gin.Context) {
	// Get the user ID from URL parameters
	id := c.Param("id")

	// Find the user by ID
	var user models.User
	if result := intializers.DB.First(&user, "id = ?", id); result.Error != nil {
		common.ReturnGenericBadResponse(c, "User not found")
		return
	}

	// Delete the user
	if result := intializers.DB.Unscoped().Delete(&user); result.Error != nil {
		common.ReturnGenericBadResponse(c, "Failed to delete user")
		return
	}

	// Return success response
	common.ReturnGenericSuccessResponse(c, "Deleted Successfully")
}

// GetAllUsers fetches all users from the database
func GetAllUsers(c *gin.Context) {
	// Declare a slice to hold the users
	var users []models.User

	// Fetch all users from the database
	if result := intializers.DB.Find(&users); result.Error != nil {
		common.ReturnGenericBadResponse(c, "Failed to fetch users")
		return
	}

	// Return the users in JSON format
	c.JSON(200, gin.H{
		"users": users,
	})
}