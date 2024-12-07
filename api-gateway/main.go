package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"time"
)

var secretKey = "yourSecretKey"

func createToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// register mock register (todo user-service)
func register(c *gin.Context) {
	var user map[string]string
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// generate JWT (TODO check in user-service)
	token, err := createToken(user["email"])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func main() {
	r := gin.Default()

	r.POST("/api/auth/register", register)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
