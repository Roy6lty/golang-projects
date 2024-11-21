package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/roy6lty/ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	//gin.Engine is a router
	// This routes do not use the middleware for authentication
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/users/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.POST("/users/productview", controllers.SearchProduct())
	incomingRoutes.POST("/users/search", controllers.SearchProductByQuery())

}
