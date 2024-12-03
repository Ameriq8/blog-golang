package routes

import (
	"encoding/json"
	"fmt"
	"go-blog/database"
	"go-blog/models"
	"regexp"

	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var body LoginBody
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if body.Phone == "" || body.Password == "" {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if !validatePhoneNumberIsIraqiPhoneNumber(body.Phone) {
		c.JSON(400, gin.H{"error": "Invalid phone number"})
		return
	}

	var user models.User
	result := database.GetDB().Where("phone = ?", body.Phone).First(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Invalid phone number"})
		return
	}

	if user.Password != body.Password {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}

	jsonData, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	// Print JSON string to console
	fmt.Println(string(jsonData))
	c.JSON(200, gin.H{"message": "Login successful"})
}

func Register(c *gin.Context) {
	var body LoginBody
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if body.Phone == "" || body.Password == "" {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if !validatePhoneNumberIsIraqiPhoneNumber(body.Phone) {
		c.JSON(400, gin.H{"error": "Invalid phone number"})
		return
	}

	if len(body.Password) < 8 {
		c.JSON(400, gin.H{"error": "Password must be at least 8 characters long"})
		return
	}

	var user models.User
	result := database.GetDB().Where("phone = ?", body.Phone).First(&user)
	if result.Error == nil {
		c.JSON(400, gin.H{"error": "Phone number already exists"})
		return
	}

	// Create new user
	newUser := models.User{
		Phone:    body.Phone,
		Password: body.Password, // Note: In production, hash this password!
	}

	result = database.GetDB().Create(&newUser)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(200, gin.H{"message": "Register successful", "user": newUser})
}

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", Login)
	router.POST("/register", Register)
}

func validatePhoneNumberIsIraqiPhoneNumber(phone string) bool {
	pattern := `^(\+964|00964|964)?(7[0-9]{8})$`
	match, _ := regexp.MatchString(pattern, phone)
	return match
}
