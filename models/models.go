package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	FirstName       *string            `json:"first_name" bson:"first_name" validate:"required,min=2,max=100"`
	LastName        *string            `json:"last_name" bson:"last_name" validate:"required,min=2,max=100"`
	Password        *string            `json:"password" bson:"password" validate:"required,min=6"`
	Email           *string            `json:"email" bson:"email" validate:"required,email"`
	Phone           *string            `json:"phone" bson:"phone" validate:"required"`
	Token           *string            `json:"token" bson:"token"`
	User_type       *string            `json:"user_type" bson:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken    *string            `json:"refresh_token" bson:"refresh_token"`
	Created_At      time.Time          `json:"created_at" bson:"created_at"`
	Updated_At      time.Time          `json:"updated_at" bson:"updated_at"`
	User_id         string             `json:"user_id" bson:"user_id"`
	UserCart        []ProductUser      `json:"product_user" bson:"product_user"`
	Address_Details []Address          `json:"address_details" bson:"address_details"`
	OrderStatus     []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	Product_Id    primitive.ObjectID `bson:"product_id"`
	Product_Name  *string            `json:"product_name" bson:"product_name"`
	Price         int64              `json:"price" bson:"price"`
	Category  	*string            `json:"category" bson:"category"`
	Image         *string            `json:"image" bson:"image"`
}

type Shop struct {
	Shop_Id   	primitive.ObjectID `bson:"shop_id"`
	Shop_Name 	*string            `json:"shop_name" bson:"shop_name"`
	Mobile    	*string            `json:"mobile" bson:"mobile"`
	Category  	*string            `json:"category" bson:"category"`
	Products 	 []Product          `json:"products" bson:"products"`
	Image     	*string            `json:"image" bson:"image"`
}

type ProductUser struct {
	Product_Id    primitive.ObjectID `bson:"product_id"`
	Product_Name  *string            `json:"product_name" bson:"product_name"`
	Price         int64              `json:"price" bson:"price"`
	Category  		*string            `json:"category" bson:"category"`
	Image         *string            `json:"image" bson:"image"`
}

type Address struct {
	Address_Id    primitive.ObjectID `bson:"address_id"`
	House         *string            `json:"house" bson:"house"`
	Street        *string            `json:"street" bson:"street"`
	City          *string            `json:"city" bson:"city"`
	Pincode       *string            `json:"pincode" bson:"pincode"`
	Lat           float64            `json:"lat" bson:"lat"`
	Long          float64            `json:"long" bson:"long"`
}

type Order struct {
	Order_Id        primitive.ObjectID `bson:"order_id"`
	Order_Cart      []ProductUser      `json:"order_cart" bson:"order_cart"`
	Ordered_At      time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price           int                `json:"price" bson:"price"`
	Discount        *int               `json:"discount" bson:"discount"`
	Payment_option  PaymentMethod      `json:"payment_option" bson:"payment_option"`
}

type PaymentMethod struct {
	Digital bool `json:"digital" bson:"digital"`
	COD     bool `json:"cod" bson:"cod"`
}