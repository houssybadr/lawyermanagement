package dto

import (
	"test/internal/models"
)

type AdminSignupRequest struct{
	User models.User `json:"user" binding:"required"`
	Admin models.Admin `json:"admin" binding:"required"`
}