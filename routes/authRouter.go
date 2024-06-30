package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mcctrix/ctrix-social-go-backend/controllers"
	"github.com/mcctrix/ctrix-social-go-backend/middleware"
)

func AuthRouter(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware)
	router.GET("login", controllers.Login())
}
