package controllers

import (
	"context"
	// "errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muhammedbilals/ecom-go/database"
	"github.com/muhammedbilals/ecom-go/helpers"
	"github.com/muhammedbilals/ecom-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var usercollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(){}

func VerifyPassword(){}

func Login() {

}

func SignUp() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx, cancel =context.WithTimeout(context.Background(),100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err!= nil {
			c.JSON(http.StatusBadRequest , gin.H{"error":err.Error()})
			return
		}

		validationeErr := validate.Struct(user)
		if validationeErr !=nil  {
			c.JSON(http.StatusBadRequest, gin.H{"error":validationeErr.Error()})
			return
		}
		count , err := usercollection.CountDocuments(ctx,bson.M{"email":user.Email})
		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error occured while checking for the email"})
		}
		count , err = usercollection.CountDocuments(ctx,bson.M{"phone":user.Phone})
		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError , gin.H{"error":"Error occured while checking for the email"})
		}
		if count>0 {
			c.JSON(http.StatusInternalServerError , gin.H{"error":"email or phone number already exist"})
			
		}
	}

}

func GetUsers(){}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		//checks if the user is admin or not
		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err := usercollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)

	}
}
