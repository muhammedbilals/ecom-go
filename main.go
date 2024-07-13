package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/muhammedbilals/ecom-go/controllers"
	"github.com/muhammedbilals/ecom-go/database"
	"github.com/muhammedbilals/ecom-go/middleware"
	"github.com/muhammedbilals/ecom-go/routes"
	// "github.com/muhammedbilals/ecom-go/tokens"
)

func main(){
	port := os.Getenv("PORT")
	if port ==""{
		port ="8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"),database.UserData(database.Client,"Users"))

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.CartCheckout())
	router.GET("/instantbuy",app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}