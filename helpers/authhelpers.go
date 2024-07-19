package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context , role string)(err error){
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("UNAUTHORIZED TO ACCESS THIS RESOURCE")
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context ,userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err =nil
	if userType =="USER" && uid !=userId{
		err =errors.New("UNAUTHORIZED TO ACCESS THIS RESOURCE")
		return err
		
	}
	err = CheckUserType(c,userType)
	return err
}