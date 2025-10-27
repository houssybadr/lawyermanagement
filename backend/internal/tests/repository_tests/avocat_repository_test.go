package tests

import(
	//"fmt"
	"test/internal/tests/testutils"
	"test/internal/models"
	"test/internal/repository"
	"testing"
)


func TestCreateAvocat(t *testing.T){
	tx:=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	authRepo:=repository.AuthRepository{}
	authRepo.SetDB(tx)

	avocatRepo:=repository.Repository[models.Avocat]{}
	avocatRepo.SetDB(tx)

	// Act
	avocatToCreate:=testutils.CreateAvocat(t,tx,authRepo)
	
	// Assert
	var fetchedAvocat models.Avocat
	if err:=avocatRepo.GetById(&fetchedAvocat,avocatToCreate.Id);err!=nil{
		t.Errorf("Failed to fetch created avocat: %v",err)
	}

	if fetchedAvocat.IsEmpty(){
		t.Errorf("Fetched avocat is empty")
	}

	if !fetchedAvocat.IsEqual(avocatToCreate){
		t.Errorf("Fetched avocat does not match created avocat")
	}
}

func TestFetchAvocatById(t *testing.T){
	tx:=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	avocatRepo:=repository.Repository[models.Avocat]{}
	avocatRepo.SetDB(tx)

	authRepo := repository.AuthRepository{}
	authRepo.SetDB(tx)
	// Act 
	avocatToCreate:=testutils.CreateAvocat(t,tx,authRepo)

	// Assert
	var fetchedAvocat models.Avocat
	if err:=avocatRepo.GetById(&fetchedAvocat,avocatToCreate.Id);err!=nil{
		t.Errorf("Failed to fetch created avocat: %v",err)
	}

	if fetchedAvocat.IsEmpty(){
		t.Errorf("Fetched avocat is empty")
	}

	if !fetchedAvocat.IsEqual(avocatToCreate){
		t.Errorf("Fetched avocat does not match created avocat")
	}
}

func TestFetchAllAvocats(t *testing.T){
	tx:=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	avocatRepo:=repository.Repository[models.Avocat]{}
	avocatRepo.SetDB(tx)
	
	authRepo := repository.AuthRepository{}
	authRepo.SetDB(tx)

	iterations:=5
	// Act
	for range iterations{
		testutils.CreateAvocat(t,tx,authRepo)
	}

	// Assert
	var avocats []models.Avocat
	if err:=avocatRepo.GetAll(&avocats);err!=nil{
		t.Errorf("Failed to fetch all avocats: %v",err)
	}

	if len(avocats)!=iterations{
		t.Errorf("Expected %d avocats, got %d",iterations,len(avocats))
	}
}

func TestUpdateAvoat(t *testing.T){
	tx:=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	avocatRepo:=repository.Repository[models.Avocat]{}
	avocatRepo.SetDB(tx)

	authRepo := repository.AuthRepository{}
	authRepo.SetDB(tx)

	// Act
	avocatToCreate:=testutils.CreateAvocat(t,tx,authRepo)
	updatedAvocatData:=testutils.GenerateRandomAvocat()
	if err:=avocatRepo.Updates(updatedAvocatData,avocatToCreate.Id);err!=nil{
		t.Errorf("Failed to update avocat: %v",err)
	}
	updatedAvocatData.AdminID=avocatToCreate.AdminID
	updatedAvocatData.UserID=avocatToCreate.UserID
	
	// Assert
	var fetchedAvocat models.Avocat
	if err:=avocatRepo.GetById(&fetchedAvocat,avocatToCreate.Id);err!=nil{
		t.Errorf("Failed to fetch updated avocat: %v",err)
	}
	if fetchedAvocat.IsEmpty(){
		t.Errorf("Fetched avocat is empty")
	}
	if !fetchedAvocat.IsEqual(updatedAvocatData){
		t.Errorf("Fetched avocat does not match updated avocat")
	}
}