package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/DahamiSathsarani/support-ticket-system/controllers"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
}
