package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "github.com/muhammedbilals/ecom-go/controllers"
	// "github.com/muhammedbilals/ecom-go/database"
	// "github.com/muhammedbilals/ecom-go/middleware"
	"github.com/muhammedbilals/ecom-go/routes"
	// "github.com/muhammedbilals/ecom-go/tokens"
)

func main(){
	err:= godotenv.Load(".env")
	if err!=nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port ==""{
		port ="8000"
	}

	// app := controllers.NewApplication(database.ProductData(database.Client, "Products"),database.UserData(database.Client,"Users"))

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.AuthRoutes(router)
	// router.Use(middleware.Authentication())

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200,gin.H{"success":"Access granded for API-1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200,gin.H{"success":"Access granded for API-2"})
	})


	// router.GET("/addtocart", app.AddToCart())
	// router.GET("/removeitem", app.RemoveItem())
	// router.GET("/cartcheckout", app.CartCheckout())
	// router.GET("/instantbuy",app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}