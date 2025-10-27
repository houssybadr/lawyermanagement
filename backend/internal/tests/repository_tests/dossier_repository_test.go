package tests

import (
	"test/internal/repository"
	"test/internal/models"
	"test/internal/tests/testutils"
	"testing"
)

func TestCreateDossier(t *testing.T){
	tx:=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	dossierRepo:=repository.Repository[models.Dossier]{}
	dossierRepo.SetDB(tx)

	// Act
	createdDossier:=testutils.CreateDossier(t,tx,dossierRepo)

	// Assert
	var fetchedDossier models.Dossier
	if err:=dossierRepo.GetById(&fetchedDossier,createdDossier.Id);err!=nil{
		t.Errorf("Failed to fetch created dossier")
	}

	if fetchedDossier.IsEmpty(){
		t.Errorf("Fetched Dossier is empty")
	}
	if !fetchedDossier.IsEqual(createdDossier){
		t.Errorf("Fetched dossier does not match created dossier")
	}
}

func TestFetchDossierByID(t *testing.T){
	tx:=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	dossierRepo:=repository.Repository[models.Dossier]{}
	dossierRepo.SetDB(tx)

	// Act
	createdDossier:=testutils.CreateDossier(t,tx,dossierRepo)
	var fetchedDossier models.Dossier
	if err:=dossierRepo.GetById(&fetchedDossier,createdDossier.Id);err!=nil{
		t.Errorf("Failed to fetch created dossier")
	}

	//Assert
	if fetchedDossier.IsEmpty(){
		t.Errorf("Fetched Dossier is empty")
	}

	if !fetchedDossier.IsEqual(createdDossier){
		t.Errorf("Fetched dossier does not match created dossier")
	}
}

func TestFetchAllDossiers(t *testing.T){
	tx:=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	dossierRepo:=repository.Repository[models.Dossier]{}
	dossierRepo.SetDB(tx)

	iterations:=5

	// Act
	for range iterations{
		testutils.CreateDossier(t,tx,dossierRepo)
	}

	//Assert
	var dossiers []models.Dossier
	if err:=dossierRepo.GetAll(&dossiers);err!=nil{
		t.Errorf("Failed to fetch all clients")
	}
	if len(dossiers)!=iterations{
		t.Errorf("Expected %d clients, got %d clients",iterations,len(dossiers))
	}
}


func TestUpdateDossier(t *testing.T){
	tx:=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	dossierRepo:=repository.Repository[models.Dossier]{}
	dossierRepo.SetDB(tx)

	//Act
	createdDossier:=testutils.CreateDossier(t,tx,dossierRepo)
	updatedDossierData:=testutils.GenerateRandomDossier()
	updatedDossierData.ClientID=createdDossier.ClientID
	if err:=dossierRepo.Updates(updatedDossierData,createdDossier.Id);err!=nil{
		t.Errorf("Failed to update client")
	}

	// Assert
	var fetchedDossier models.Dossier
	if err:=dossierRepo.GetById(&fetchedDossier,createdDossier.Id);err!=nil{
		t.Errorf("Failed to fetch updated dossier")
	}

	if fetchedDossier.IsEmpty(){
		t.Errorf("Fetched dossier is Empty")
	}
	if !fetchedDossier.IsEqual(updatedDossierData){
		t.Errorf("Fetched client does not match created dossier")
	}
}

func TestDeleteDossier(t *testing.T){
	tx :=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	dossierRepo:=repository.Repository[models.Dossier]{}
	dossierRepo.SetDB(tx)

	//Act 
	dossierToDelete:=testutils.CreateDossier(t,tx,dossierRepo)
	if err:=dossierRepo.Delete(dossierToDelete.Id);err!=nil{
		t.Errorf("Failed to delete dossier")
	}

	// Assert
	var dossierCount int64
	dossierRepo.Count(&dossierCount)
	if dossierCount!=0{
		t.Errorf("Failed to delete dossier")
	}

}