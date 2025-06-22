package controllers

import (
	"net/http"
	"github.com/DahamiSathsarani/support-ticket-system/database"
	"github.com/DahamiSathsarani/support-ticket-system/models"

	"github.com/gin-gonic/gin"
)

func CreateTicket(c *gin.Context) {
	var ticket models.Ticket

	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("userID")
	ticket.UserID = userID

	if err := database.DB.Create(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ticket"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func GetUserTickets(c *gin.Context) {
	userID := c.GetUint("userID")

	var tickets []models.Ticket
	database.DB.Where("user_id = ?", userID).Find(&tickets)

	c.JSON(http.StatusOK, tickets) 
}

func GetAllTickets(c *gin.Context) {
	var tickets []models.Ticket
	database.DB.Find(&tickets)

	c.JSON(http.StatusOK, tickets)
}

func UpdateTicketStatus(c *gin.Context) {
	var ticket models.Ticket
	id := c.Param("id")

	if err := database.DB.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&ticket)
	c.JSON(http.StatusOK, ticket)
}

func DeleteTicket(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Ticket{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted"})
}
