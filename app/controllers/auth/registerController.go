package auth

import (
	"log"
	"net/http"

	"github.com/arifikhsan/gin_rest/app/models"
	db "github.com/arifikhsan/gin_rest/databases"
	"github.com/arifikhsan/gin_rest/libraries/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	user.Password = string(hash)
	if err = db.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "status": http.StatusUnprocessableEntity})
		return
	}
	accessToken := jwt.GenerateToken(user.ID)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "messages": "Sucessfully registered", "accessToken": accessToken})
}
