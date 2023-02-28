package routes

import (
	controller "github.com/dasarathg68/GoLangAuthJWT/controllers"
	"github.com/dasarathg68/GoLangAuthJWT/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
