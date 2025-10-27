package tests

import (
	"testing"
	"test/internal/models"
	"test/internal/repository"
	"test/internal/tests/testutils"

	//"fmt"
)

func TestCreateClient(t *testing.T){
	tx :=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	clientRepo:=repository.Repository[models.Client]{}
	clientRepo.SetDB(tx)

	// Act
	clientToCreate:=testutils.CreateClient(t,tx,clientRepo)

	// Assert
	var fetchedClient models.Client
	if err:=clientRepo.GetById(&fetchedClient,clientToCreate.Id);err!=nil{
		t.Errorf("Failed To Fetch created client %v",err)
	}

	if fetchedClient.IsEmpty(){
		t.Errorf("Fetched client is empty")
	}
	
	if !fetchedClient.IsEqual(clientToCreate){
		t.Errorf("Fetched client does not match created client")
	}
}

func TestFetchClientById(t *testing.T){
	tx :=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	clientRepo:=repository.Repository[models.Client]{}
	clientRepo.SetDB(tx)

	// Act
	clientToCreate:=testutils.CreateClient(t,tx,clientRepo)
	var fetchedClient models.Client
	if err:=clientRepo.GetById(&fetchedClient,clientToCreate.Id);err!=nil{
		t.Errorf("Failed to fetch created client %v",err)
	}

	// Assert
	if fetchedClient.IsEmpty(){
		t.Errorf("Fetched client is Empty")
	}

	if !fetchedClient.IsEqual(clientToCreate){
		t.Errorf("Fetched client does not match created client")
	}
}

func TestFetchAllClients(t *testing.T){
	tx :=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	clientRepo:=repository.Repository[models.Client]{}
	clientRepo.SetDB(tx)

	iterations:=5

	// Act
	for range iterations{
		testutils.CreateClient(t,tx,clientRepo)
	}

	//Assert
	var fetchedClients []models.Client
	if err:=clientRepo.GetAll(&fetchedClients);err!=nil{
		t.Errorf("Failed to fetch all clients")
	}
	if len(fetchedClients)!=iterations{
		t.Errorf("Expected %d clients, got %d clients",iterations,len(fetchedClients))
	}
}

func TestUpdateClient(t *testing.T){
	tx :=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	clientRepo:=repository.Repository[models.Client]{}
	clientRepo.SetDB(tx)

	//Act 
	clientToCreate:=testutils.CreateClient(t,tx,clientRepo)
	updatedClientData:=testutils.GenerateRandomClient()
	updatedClientData.AvocatID=clientToCreate.AvocatID
	if err:=clientRepo.Updates(updatedClientData,clientToCreate.Id);err!=nil{
		t.Errorf("Failed to update client")
	}

	// Assert
	var fetchedClient models.Client
	if err:=clientRepo.GetById(&fetchedClient,clientToCreate.Id);err!=nil{
		t.Errorf("Failed to fetch updated client")
	}

	if fetchedClient.IsEmpty(){
		t.Errorf("Fetched client is Empty")
	}
	if !fetchedClient.IsEqual(updatedClientData){
		t.Errorf("Fetched client does not match created client")
	}
}

func TestDeleteClient(t *testing.T){
	tx :=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	clientRepo:=repository.Repository[models.Client]{}
	clientRepo.SetDB(tx)

	//Act 
	clientToDelete:=testutils.CreateClient(t,tx,clientRepo)
	if err:=clientRepo.Delete(clientToDelete.Id);err!=nil{
		t.Errorf("Failed to delete client")
	}

	// Assert
	var clientCount int64
	clientRepo.Count(&clientCount)
	if clientCount!=0{
		t.Errorf("Failed to delete client")
	}

}