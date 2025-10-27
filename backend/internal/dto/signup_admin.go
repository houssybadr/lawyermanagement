package dto

import (
	"github.com/houssybadr/lawyermanagement/backend/internal/models"
)

type AdminSignupRequest struct{
	User models.User `json:"user" binding:"required"`
	Admin models.Admin `json:"admin" binding:"required"`
}