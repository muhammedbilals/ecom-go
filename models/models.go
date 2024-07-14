package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID				primitive.ObjectID	`bson:"_id"`
	FirstName 		*string				 `json:"first_name" validate:"required,min=2,max=100"`
	LastName		*string				 `json:"last_name" validate:"required,min=2,max=100"`
	Password		*string              `json:"password" validate:"required,min=6"`
	Email			*string      	     `json:"email" validate:"email, requied"`
	Phone			*string				 `json:"phone" validate:"phone, requied"`
	Token			*string			 	 `json:"token"`
	User_type		*string				 `json:"user_type" validate:"requied ,eq=ADMIN|eq=USER"`
	Refresh_token	*string				 `json:"refresh_token"`
	Created_at		time.Time			 `json:"created_at"`
	Updated_at		time.Time			 `json:"updated_at"`
	User_id			*string				 `json:"user_id"`

}