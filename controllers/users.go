package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/muhammedbilals/ecom-go/helpers"
	"github.com/muhammedbilals/ecom-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		//checking user type ,grands access if only ADMIN to this endpoint
		// if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

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

		cursor, err1 := usercollection.Aggregate(ctx, pipeline)
		if err1 != nil {
            c.JSON(500, gin.H{"error": err1.Error()})
            return
        }

		defer cursor.Close(ctx)


		var results []bson.M
        if err = cursor.All(ctx, &results); err != nil {
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
