package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DahamiSathsarani/support-ticket-system/database"
	"github.com/DahamiSathsarani/support-ticket-system/models"
	"github.com/DahamiSathsarani/support-ticket-system/routes"
)

func main() {
	r := gin.Default()

	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	routes.AuthRoutes(r)

	r.Run(":8080")
}
