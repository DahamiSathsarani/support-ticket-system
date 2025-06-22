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

func UpdateTicket(c *gin.Context) {
	var ticket models.Ticket
	id := c.Param("id")
	userID := c.GetUint("userID")
	role := c.GetString("role")

	if err := database.DB.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	if ticket.UserID != userID && role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.Ticket
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket.Title = input.Title
	ticket.Description = input.Description

	if role == "admin" {
		ticket.Status = input.Status
	}

	if err := database.DB.Save(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ticket"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func DeleteTicket(c *gin.Context) {
	var ticket models.Ticket
	id := c.Param("id")
	userID := c.GetUint("userID")
	role := c.GetString("role")

	if err := database.DB.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	if ticket.UserID != userID && role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	database.DB.Delete(&ticket)
	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted"})
}
