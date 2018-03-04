package middlewares

import (
	"net/http"

	"github.com/arifikhsan/gin_rest/libraries/jwt"
	"github.com/gin-gonic/gin"
)

func Authorized(c *gin.Context) {
	unparsedToken := c.Request.Header.Get("Authorization")
	if unparsedToken == "" {
		c.JSON(http.StatusForbidden, gin.H{"messages": "Authorization token required", "status": http.StatusForbidden})
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	granted, messages := jwt.CheckToken(unparsedToken)
	if !granted {
		c.JSON(http.StatusUnauthorized, gin.H{"messages": messages, "status": http.StatusUnauthorized})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
