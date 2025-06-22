package routes

import (
	"github.com/DahamiSathsarani/support-ticket-system/controllers"
	"github.com/DahamiSathsarani/support-ticket-system/middleware"

	"github.com/gin-gonic/gin"
)

func TicketRoutes(router *gin.Engine) {
	ticket := router.Group("/api/tickets")
	ticket.Use(middleware.AuthMiddleware()) 

	ticket.POST("/create", controllers.CreateTicket)
	ticket.GET("/get-my", controllers.GetUserTickets)

	ticket.GET("/get-all", middleware.RoleMiddleware("admin"), controllers.GetAllTickets)

	ticket.PUT("/update/:id", controllers.UpdateTicket)
	ticket.DELETE("/delete/:id", controllers.DeleteTicket)
}
