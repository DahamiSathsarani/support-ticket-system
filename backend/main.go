package main

import (
	"log"
	"time"
	"os"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/DahamiSathsarani/support-ticket-system/backend/database"
	"github.com/DahamiSathsarani/support-ticket-system/backend/models"
	"github.com/DahamiSathsarani/support-ticket-system/backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		log.Fatal("FRONTEND_URL is not set in .env")
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.Ticket{})

	routes.AuthRoutes(r)
	routes.TicketRoutes(r)
	routes.UserRoutes(r)


	r.Run(":8080")
}
