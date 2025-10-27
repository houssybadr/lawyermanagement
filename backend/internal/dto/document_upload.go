package dto

import (
	"time"
)

type DocumentUploadResponse struct{
	Id           uint      
	Nom          string    
	DateCreation time.Time 
	DossierID    uint      
}