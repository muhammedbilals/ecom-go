package middleware

import (

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedbilals/ecom-go/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clienToken := c.Request.Header.Get("token")	
		if clienToken == ""{
			c.JSON(http.StatusInternalServerError ,gin.H{"error":"No Autherization header provided"})
			c.Abort()
			return
		}

		claims ,err :=helpers.ValidateToken(clienToken)
		if err != ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err})
			c.Abort()
			return
		}
		c.Set("email" ,claims.Email)
		c.Set("first_name" ,claims.FirstName)
		c.Set("last_name" ,claims.LastName)
		c.Set("uid" ,claims.Uid)
		c.Set("user_type" ,claims.User_type)
		c.Next()
	}
}