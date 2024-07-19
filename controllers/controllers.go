package controllers

import (
	"context"
	"fmt"
	"strconv"
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muhammedbilals/ecom-go/database"
	"github.com/muhammedbilals/ecom-go/helpers"
	"github.com/muhammedbilals/ecom-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var usercollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic()
	}
	return string(bytes)
}

func VerifyPassword(userpassword string, providepassword string) (bool, string) {
	//compare password with bycrypt
	err := bcrypt.CompareHashAndPassword([]byte(providepassword), []byte(userpassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprint("email or password is incorret")
	}
	return check, msg
}



func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		var foundUser models.User

		//Get user from api call
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error : ": err.Error()})
		}
		//checks if the user is found on database and passing it to foundUser variable
		err := usercollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}
		//verify the password with bcrypt
		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if !passwordIsValid  {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		}
		if foundUser.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		}
		token, refereshToken, err := helpers.GenerateAlltokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, *foundUser.User_type, foundUser.User_id)
		helpers.UpdateAllTokens(token, refereshToken, foundUser.User_id)

		usercollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, foundUser)
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationeErr := validate.Struct(user)
		if validationeErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationeErr.Error()})
			return
		}
		count, err := usercollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking for the email"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password
		count, err = usercollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking for the email"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or phone number already exist"})
			return
		}
		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, err := helpers.GenerateAlltokens(*user.Email, *user.FirstName, *user.LastName, *user.User_type, user.User_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
			return
		}
		user.Token = &token
		user.RefreshToken = &refreshToken

		//insert into databsse
		resultInsertNumber, InsertError := usercollection.InsertOne(ctx, user)
		if InsertError != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertNumber)
	}

}

func GetUsers() gin.HandlerFunc{
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c ,"ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest , gin.H{"error":err.Error()})
		}
		var ctx , cancel  =context.WithTimeout(context.Background(),100*time.Second)

		recordPerPage , err:= strconv.Atoi(c.Query("recordPerPage"))

		if err!=nil || recordPerPage<1{
			recordPerPage =0
		}
		page , err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page<1{
			page =1
		}
		startIndex := (page -1 )*recordPerPage
		startIndex , err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
		groupStage := bson.D{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{{Key: "_id",Value: "null"}}},
			{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
			{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
		}}}
		projectStage := bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "total_count", Value: 1},
			{Key: "user_items", Value: bson.D{ 
				{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}},
			}},
		}}}

		result, err:=usercollection.Aggregate(ctx,mongo.Pipeline{
			matchStage, groupStage , projectStage,
		})

		defer cancel()

		if err!=nil{
			c.JSON(http.StatusInternalServerError ,gin.H{"error":"error occured while lising user items"})
		}
		var allUsers []bson.M
		if err := result.All(ctx, &allUsers); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while decoding user data"})
			return
		}
		c.JSON(http.StatusOK, allUsers[0])
	}
}

func GetUser() gin.HandlerFunc {
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
