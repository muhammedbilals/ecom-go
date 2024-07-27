package controllers

// import (
// 	"context"
// 	"errors"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/muhammedbilals/ecom-go/database"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )


// type Application struct{
// 	prodCollection *mongo.Collection
// 	usercollection *mongo.Collection
// }

// func NewApplicaiton(prodCollection *mongo.Collection, usercollection *mongo.Collection)*Application{
// 	return &Application{
// 		prodCollection: prodCollection,
// 		usercollection: usercollection,
// 	}
// }

// func (app *Application) AddToCart() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		productQuieryId := c.Query("id") 
// 		if productQuieryId==""{
// 			log.Println("product id is empty")

// 			_ = c.AbortWithError(http.StatusBadRequest,errors.New("product id is empty"))
// 			return
// 		}

// 		userQuieryId := c.Query("userId") 
// 		if userQuieryId==""{
// 			log.Println("product id is empty")

// 			_ = c.AbortWithError(http.StatusBadRequest,errors.New("product id is empty"))
// 			return
// 		}

// 		productId ,err := primitive.ObjectIDFromHex(productQuieryId)
// 		if err!=nil{
// 			log.Println(err)
// 			c.AbortWithStatus(http.StatusInternalServerError)
// 			return
// 		}
		
// 		var ctx ,cancel = context.WithTimeout(context.Background(), time.Second*100)
// 		defer cancel()

// 		err = database.AddProductToCart(ctx, app.prodCollection,app.usercollection,productId,userQuieryId)
// 		if err!=nil{
// 			c.IndentedJSON(http.StatusInternalServerError,err)

// 		}
// 		c.IndentedJSON(200,"Successfully added to cart")
// 	}
// }

// func (app *Application) RemoveItem() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		productQuieryId := c.Query("id") 
// 		if productQuieryId==""{
// 			log.Println("product id is empty")

// 			_ = c.AbortWithError(http.StatusBadRequest,errors.New("product id is empty"))
// 			return
// 		}

// 		userQuieryId := c.Query("userId") 
// 		if userQuieryId==""{
// 			log.Println("product id is empty")

// 			_ = c.AbortWithError(http.StatusBadRequest,errors.New("product id is empty"))
// 			return
// 		}

// 		productId ,err := primitive.ObjectIDFromHex(productQuieryId)
// 		if err!=nil{
// 			log.Println(err)
// 			c.AbortWithStatus(http.StatusInternalServerError)
// 			return
// 		}
		
// 		var ctx ,cancel = context.WithTimeout(context.Background(), time.Second*100)
// 		defer cancel()

// 		err = database.RemoveCartItem(ctx, app.prodCollection,app.usercollection,productId,userQuieryId)
// 		if err!=nil{
// 			c.IndentedJSON(http.StatusInternalServerError,err)
// 		}
// 		c.IndentedJSON(200,"Successfully added to cart")
// 	}
// }

// func (app *Application) BuyItemFromCart() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		productQuieryId := c.Query("id") 
// 		if productQuieryId==""{
// 			log.Println("product id is empty")

// 			_ = c.AbortWithError(http.StatusBadRequest,errors.New("product id is empty"))
// 			return
// 		}

// 		userQuieryId := c.Query("userId") 
// 		if userQuieryId==""{
// 			log.Println("product id is empty")

// 			_ = c.AbortWithError(http.StatusBadRequest,errors.New("product id is empty"))
// 			return
// 		}

		
// 		var ctx ,cancel = context.WithTimeout(context.Background(), time.Second*100)
// 		defer cancel()

// 		// err = database.BuyItemFromCart(ctx, app.prodCollection,app.usercollection,productId,userQuieryId)
// 		// if err!=nil{
// 		// 	c.IndentedJSON(http.StatusInternalServerError,err)
// 		// }
// 		c.IndentedJSON(200,"Successfully added to cart")
// 	}
// }

