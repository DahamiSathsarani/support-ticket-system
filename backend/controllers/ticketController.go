package controllers

import (
	"net/http"
	"github.com/DahamiSathsarani/support-ticket-system/backend/database"
	"github.com/DahamiSathsarani/support-ticket-system/backend/models"

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

func AssignTicket(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can assign tickets"})
		return
	}

	var ticket models.Ticket
	id := c.Param("id")

	if err := database.DB.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	var input struct {
		AssignedTo uint `json:"assigned_to"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var assignedUser models.User
	if err := database.DB.First(&assignedUser, input.AssignedTo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Assigned user not found"})
		return
	}

	if assignedUser.Role != "agent" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only users with 'agent' role can be assigned tickets"})
		return
	}

	ticket.AssignedTo = input.AssignedTo

	if err := database.DB.Save(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket assigned successfully",
		"ticket":  ticket,
	})
}

func GetAssignedTickets(c *gin.Context) {
	userID := c.GetUint("userID")

	var tickets []models.Ticket
	if err := database.DB.Where("assigned_to = ?", userID).Find(&tickets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch assigned tickets"})
		return
	}

	if len(tickets) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No assigned tickets"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func GetUserTicketStats(c *gin.Context) {
	userID := c.GetUint("userID")

	var total, open, pending, resolved, rejected int64

	database.DB.Model(&models.Ticket{}).Where("user_id = ?", userID).Count(&total)
	database.DB.Model(&models.Ticket{}).Where("user_id = ? AND status = ?", userID, "open").Count(&open)
	database.DB.Model(&models.Ticket{}).Where("user_id = ? AND status = ?", userID, "pending").Count(&pending)
	database.DB.Model(&models.Ticket{}).Where("user_id = ? AND status = ?", userID, "resolved").Count(&resolved)
	database.DB.Model(&models.Ticket{}).Where("user_id = ? AND status = ?", userID, "rejected").Count(&rejected)

	c.JSON(http.StatusOK, gin.H{
		"total":    total,
		"open":     open,
		"pending":  pending,
		"resolved": resolved,
		"rejected": rejected,
	})
}
