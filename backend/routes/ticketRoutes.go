package routes

import (
	"github.com/DahamiSathsarani/support-ticket-system/backend/controllers"
	"github.com/DahamiSathsarani/support-ticket-system/backend/middleware"

	"github.com/gin-gonic/gin"
)

func TicketRoutes(router *gin.Engine) {
	ticket := router.Group("/api/tickets")
	ticket.Use(middleware.AuthMiddleware()) 

	ticket.POST("/create", controllers.CreateTicket)
	ticket.GET("/get-my", controllers.GetUserTickets)
	ticket.PUT("/update/:id", controllers.UpdateTicket)
	ticket.DELETE("/delete/:id", controllers.DeleteTicket)
	ticket.GET("/get-assigned", controllers.GetAssignedTickets)
	ticket.GET("/get-tickets-stats", controllers.GetUserTicketStats)

	ticket.GET("/get-all", middleware.RoleMiddleware("admin"), controllers.GetAllTickets)
	ticket.PUT("/assign/:id", middleware.RoleMiddleware("admin"), controllers.AssignTicket)

}
