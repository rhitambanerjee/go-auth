package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rhitambanerjee/GOlang-JWT-Project/controllers"
	middleware "github.com/rhitambanerjee/GOlang-JWT-Project/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:id", controller.GetUser())
}