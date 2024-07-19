package helpers

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/muhammedbilals/ecom-go/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	LastName  string
	FirstName string

	Email     string
	Uid       string
	User_type string
	jwt.StandardClaims
	User_id string
}

var usercollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateAlltokens(email string, firstName string, lastName string, userType string, uid string) (string, string, error) {
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Uid:       uid,
		User_type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}
	return token, refreshToken, err
}
