package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SecretRoom(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"messages": "sekret rum", "status": http.StatusOK})
}
