package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/muhammedbilals/ecom-go/database"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

type SignedDetails struct{
	
	LastName		string				
	FirstName 		string			
	
	Email			string      	   
	Uid 			string
	User_type		string			
	jwt.StandardClaims
	User_id			string			
}

var usercollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY = os.Getenv("SECRET_KEY")