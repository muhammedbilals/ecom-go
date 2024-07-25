package helpers

import (
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	uid := c.GetString("user_id")
	err =nil
	if userType =="USER" && uid !=userId{
		err =errors.New("UNAUTHORIZED TO ACCESS THIS RESOURCE")
		return err
		
	}
	err = CheckUserType(c,userType)
	return err
}

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