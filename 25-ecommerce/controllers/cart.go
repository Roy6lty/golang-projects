package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/roy6lty/ecommerce/models"

	"github.com/gin-gonic/gin"
	"github.com/roy6lty/ecommerce/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection *mongo.Collection, userCollection *mongo.Collection) *Application {
	//A mongodb collection will be passed into the func
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id") //get the queryID from gin context
		if productQueryID == "" {
			log.Println("Product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product id is empty")) //returns a response with an error
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("User id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("User id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(http.StatusOK, "Successfully added to cart")

	}

}

func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id") //get the queryID from gin context
		if productQueryID == "" {
			log.Println("Product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product id is empty")) //returns a response with an error
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("User id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("User id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.IndentedJSON(200, "Successfully removed item from cart")

	}

}

func GetItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid search id"})
			c.Abort()
			return

		}

		usert_id, _ := primitive.ObjectIDFromHex(user_id)

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var filledCart models.User
		err := UserCollection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: usert_id}}).Decode(&filledCart)

		if err != nil {
			log.Println(err)
			c.IndentedJSON(500, "not found")
			return
		}

		//Key is a mongodb Keyword and value is the data you want this operation performed on

		filter_match := bson.D{{Key: "Search", Value: bson.D{primitive.E{Key: "_id", Value: usert_id}}}}
		unwind := bson.D{{Key: "$Unwind", Value: bson.D{primitive.E{Key: "path", Value: "$userCart"}}}}
		grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{primitive.E{Key: "$sum", Value: "$userCart.price"}}}}}}
		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline(filter_match, unwind, grouping))

		if err != nil {
			log.Println(err)
		}
		var listing []bson.M
		if err = pointCursor.All(ctx, &listing); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		for _, json := range listing {
			c.IndentedJSON(200, json["total"])
			c.IndentedJSON(200, filledCart.UserCart)
		}
		ctx.Done()
	}
}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")

		if userQueryID == "" {
			log.Println("UserId is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("UserID is empty"))
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		err := database.BuyItemFromCart(ctx, app.userCollection, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)

		}
		c.IndentedJSON("Successfully placed an order")

	}

}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id") //get the queryID from gin context
		if productQueryID == "" {
			log.Println("Product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product id is empty")) //returns a response with an error
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("User id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("User id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.InstantBuyer(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

	}

}
