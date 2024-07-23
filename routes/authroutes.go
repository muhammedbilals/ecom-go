package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedbilals/ecom-go/controllers"
	
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.GET("/users/login", controllers.Login())
}
