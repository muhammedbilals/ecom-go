package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


type Application struct{
	prodCollection *mongo.Collection
	usercollection *mongo.Collection
}

func NewApplicaiton(prodCollection *mongo.Collection, usercollection *mongo.Collection)*Application{
	return &Application{
		prodCollection: prodCollection,
		usercollection: usercollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}