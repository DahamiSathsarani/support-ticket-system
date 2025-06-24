package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/DahamiSathsarani/support-ticket-system/middleware"
	"github.com/DahamiSathsarani/support-ticket-system/controllers"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/api/users")
	user.Use(middleware.AuthMiddleware())

	user.GET("/get", controllers.GetUser)
	user.PUT("/update-password", controllers.UpdatePassword)
}
