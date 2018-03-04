package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/arifikhsan/gin_rest/app/controllers"
	"github.com/arifikhsan/gin_rest/app/controllers/auth"
	"github.com/arifikhsan/gin_rest/app/controllers/middlewares"
)

func init() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	authorized := r.Group("/")
	authorized.Use(middlewares.Authorized)
	{
		private := authorized.Group("api/v1/")
		private.GET("/secretroom", controllers.SecretRoom)
		private.POST("person", controllers.CreatePerson)
		private.GET("person", controllers.GetAllPerson)
		private.GET("person/:id", controllers.GetOnePerson)
		private.PUT("person/:id", controllers.UpdatePerson)
		private.DELETE("person/:id", controllers.DeletePerson)
	}

	unAuthorized := r.Group("/api/v1/")
	{
		unAuthorized.POST("/register", auth.Register)
		unAuthorized.POST("/login", auth.Login)
	}

	r.Run(":8080")
}
