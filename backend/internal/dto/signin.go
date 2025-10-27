package dto

import (
	"test/internal/models"
)

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SuccessfulAdminSignInResponse struct {
	Token string       `json:"token" binding:"required"`
	Role  models.Role  `json:"role" binding:"required"`
	Admin models.Admin `json:"admin" binding:"required"`
}

type SuccessfulAvocatSignInResponse struct {
	Token  string        `json:"token" binding:"required"`
	Role   models.Role   `json:"role" binding:"required"`
	Avocat models.Avocat `json:"admin" binding:"required"`
}
