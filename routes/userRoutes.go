package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedbilals/ecom-go/controllers"
	"github.com/muhammedbilals/ecom-go/middleware"
)



func UserRoutes(incomingRouts * gin.Engine){
	incomingRouts.Use(middleware.Authenticate())
	incomingRouts.GET("/users",controllers.GetUsers())
	incomingRouts.GET("/users/:user_id",controllers.GetUser())

	// incomingRouts.POST("/users/signup , ",controllers.SignUp())
	// incomingRouts.POST("/users/login , ",controllers.Login())
	// incomingRouts.POST("/admin/addproduct",controllers.AddProduct())
	// incomingRouts.GET("/users/productview , ",controllers.ProductView)
	// incomingRouts.GET("/users/search , ",controllers.Search())
}

