package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/DahamiSathsarani/support-ticket-system/middleware"
)

func ProtectedRoutes(router *gin.Engine) {
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/profile", func(c *gin.Context) {
		userID := c.GetUint("userID")
		role := c.GetString("role")

		c.JSON(200, gin.H{
			"message": "You are logged in!",
			"userID":  userID,
			"role":    role,
		})
	})
}
