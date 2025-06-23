package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status" gorm:"default:open"` 
	UserID      uint   `json:"user_id"`         
	AssignedTo  uint   `json:"assigned_to"`           
}
