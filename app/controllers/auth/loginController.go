package auth

import (
	"net/http"

	"github.com/arifikhsan/gin_rest/app/models"
	db "github.com/arifikhsan/gin_rest/databases"
	"github.com/arifikhsan/gin_rest/libraries/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	inputPassword := user.Password
	if err := db.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": http.StatusUnauthorized})
		return
	}
	passwordMatch := CheckPasswordHash(inputPassword, user.Password)
	if !passwordMatch {
		c.JSON(http.StatusUnauthorized, gin.H{"messages": "Wrong password", "status": http.StatusUnauthorized})
		return
	}

	accessToken := jwt.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "messages": "Logged in....", "status": http.StatusOK})
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
