package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/DahamiSathsarani/support-ticket-system/backend/database"
	"github.com/DahamiSathsarani/support-ticket-system/backend/models"
	"github.com/DahamiSathsarani/support-ticket-system/backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.Ticket{})

	routes.AuthRoutes(r)
	routes.TicketRoutes(r)
	routes.UserRoutes(r)


	r.Run(":8080")
}
