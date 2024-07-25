package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

var usercollection *mongo.Collection = database.OpenCollection(database.Client, "users")
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

		if c.Request.Method != http.MethodGet {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

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
		if !passwordIsValid {
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		}
		if foundUser.Email == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		}
		token, refereshToken, err := helpers.GenerateAlltokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, *foundUser.User_type, foundUser.User_id)
		helpers.UpdateAllTokens(token, refereshToken, foundUser.User_id)

		//finding user with the user_id
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
		//getting total document count
		count, err := usercollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking for the email"})
			return
		}

		//hashing password with the helper function
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
		//adding time
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

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		//checking user type ,grands access if only ADMIN to this endpoint
		// if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// defer cancel()

		//converting string to intiger with strconv 
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		//returns 10 users by default
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		//gets the page no from url
		page, err1 := strconv.Atoi(c.Query("page"))
		//set defalut page to 1
		if err1 != nil || page < 1 {
			page = 1
		}

		//
		startIndex := (page - 1) * recordPerPage
		// startIndex, err = strconv.Atoi(c.Query("startIndex"))

		
		pipeline := mongo.Pipeline{
            {{Key: "$match", Value: bson.D{{Key: "user_type", Value: "ADMIN"}}}}, // Filter for user_type "ADMIN"
            {{Key: "$project", Value: bson.D{ // Project specific fields
                {Key: "firstname", Value: 1},
                {Key: "lastname", Value: 1},
                {Key: "email", Value: 1},
                {Key: "phone", Value: 1},
                {Key: "created_at", Value: 1},
                {Key: "updated_at", Value: 1},
            }}},
            {{Key: "$sort", Value: bson.D{{Key: "created_at", Value: -1}}}}, // Sort by creation date in descending order
            {{Key: "$skip", Value: startIndex}},                // Skip documents based on startIndex
            {{Key: "$limit", Value: recordPerPage}},            // Limit the number of documents
        }

		cursor, err1 := usercollection.Aggregate(context.TODO(), pipeline)
		if err1 != nil {
            c.JSON(500, gin.H{"error": err1.Error()})
            return
        }

		defer cursor.Close(context.TODO())


		var results []bson.M
        if err = cursor.All(context.TODO(), &results); err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
		c.JSON(200, results)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		//checks if the user is admin or not
		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		//getting the user with the user_id
		err := usercollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)

	}
}
