package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/roy6lty/ecommerce/controllers"
	"github.com/roy6lty/ecommerce/database"
	"github.com/roy6lty/ecommerce/middleware"
)

func main() {
	port := os.Gotenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	router.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
