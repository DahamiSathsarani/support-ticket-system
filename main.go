package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/DahamiSathsarani/support-ticket-system/database"
	"github.com/DahamiSathsarani/support-ticket-system/models"
	"github.com/DahamiSathsarani/support-ticket-system/routes"
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
	routes.ProtectedRoutes(r)

	r.Run(":8080")
}
