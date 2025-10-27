package testutils

import (
	"test/internal/models"
	"test/internal/repository"
	"testing"

	"gorm.io/gorm"
)

func CreateAdmin(t *testing.T,tx *gorm.DB,authRepo repository.AuthRepository) models.Admin{
	userToCreate := GenerateRandomUser()
	adminToCreate :=GenerateRandomAdmin()
	userToCreate.Role = models.AdminRole

	if err := authRepo.SignUp(&userToCreate, &adminToCreate); err != nil {
		t.Errorf("Failed to create admin: %v", err)
	}

	return adminToCreate
}

func CreateAvocat(t *testing.T,tx *gorm.DB,authRepo repository.AuthRepository) models.Avocat{
	userToCreateForAvocat:=GenerateRandomUser()
	avocatToCreate:=GenerateRandomAvocat()
	createdAdmin:=CreateAdmin(t,tx,authRepo)
	
	avocatToCreate.AdminID=createdAdmin.Id
	if err:=authRepo.SignUp(&userToCreateForAvocat,&avocatToCreate);err!=nil{
		t.Errorf("Failed to create avocat %v",err)
	}

	return avocatToCreate
}

func CreateClient(t *testing.T,tx *gorm.DB,clientRepo repository.Repository[models.Client]) models.Client{
	authRepo := repository.AuthRepository{}
	authRepo.SetDB(tx)

	createdAvocat:=CreateAvocat(t,tx,authRepo)
	clientToCreate:=GenerateRandomClient()
	
	clientToCreate.AvocatID=createdAvocat.Id
	if err:=clientRepo.Create(&clientToCreate);err!=nil{
		t.Errorf("Failed to create client %v",err)
	}

	return clientToCreate
}

func CreateDossier(t *testing.T,tx *gorm.DB,dossierRepo repository.Repository[models.Dossier]) models.Dossier{
	clientRepo:=repository.Repository[models.Client]{}
	clientRepo.SetDB(tx)

	createdClient:=CreateClient(t,tx,clientRepo)
	dossierToCreate:=GenerateRandomDossier()
	
	dossierToCreate.ClientID=createdClient.Id
	if err:=dossierRepo.Create(&dossierToCreate);err!=nil{
		t.Errorf("Failed to create dossier %v",err)
	}
	return dossierToCreate
}

func CreateDocument(t *testing.T,tx *gorm.DB,documentRepository repository.Repository[models.Document])models.Document{
	dossierRepository:= repository.Repository[models.Dossier]{}
	dossierRepository.SetDB(tx)
	
	createdDossier:=CreateDossier(t,tx,dossierRepository)
	documerntToCreate:=GenerateRandomDocument()
	documerntToCreate.DossierID=createdDossier.Id

	if err:=documentRepository.Create(&documerntToCreate);err!=nil{
		t.Error("Failed to create document")
	}

	return documerntToCreate
}