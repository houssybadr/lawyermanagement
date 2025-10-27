package dto 

import(
	"test/internal/models"
)

type AvocatSignupRequest struct{
	User models.User `json:"user" binding:"required"`
	Avocat models.Avocat `json:"avocat" binding:"required"`
}