package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedbilals/ecom-go/controllers"
)

func AuthRoutes(incomingRouts *gin.Engine) {
	incomingRouts.POST("/users/signup , ", controllers.SignUp())
	incomingRouts.GET("/users/login , ", controllers.Login())
}
