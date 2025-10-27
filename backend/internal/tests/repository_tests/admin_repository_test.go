package tests

import (
	//"fmt"
	"test/internal/models"
	"test/internal/repository"
	"test/internal/tests/testutils"
	"testing"
)

func TestCreateAdmin(t *testing.T) {
	tx := testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	adminRepo := repository.Repository[models.Admin]{}
	adminRepo.SetDB(tx)

	authRepo := repository.AuthRepository{}
	authRepo.SetDB(tx)
	// Act
	adminToCreate := testutils.CreateAdmin(t,tx,authRepo)

	// Assert
	var fetcheAdmin models.Admin
	if err := adminRepo.GetById(&fetcheAdmin, adminToCreate.Id); err != nil {
		t.Errorf("Failed to fetch created admin: %v", err)
	}

	if fetcheAdmin.IsEmpty() {
		t.Errorf("Fetched admin is empty")
	}

	if !fetcheAdmin.IsEqual(adminToCreate) {
		t.Errorf("Fetched admin does not match created admin")
	}

}

func TestFecthAdminById(t *testing.T) {
	tx := testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	authRepo := repository.AuthRepository{}
	authRepo.SetDB(tx)

	adminRepo := repository.Repository[models.Admin]{}
	adminRepo.SetDB(tx)

	// Act
	adminToCreate:= testutils.CreateAdmin(t,tx,authRepo)

	// Assert
	var fetchedAdmin models.Admin
	if err:=adminRepo.GetById(&fetchedAdmin,adminToCreate.Id);err!=nil{
		t.Errorf("Failed to fetch created admin: %v",err)
	}
	if fetchedAdmin.IsEmpty() {
		t.Errorf("Fetched admin is empty")
	}

	if !fetchedAdmin.IsEqual(adminToCreate) {
		t.Errorf("Fetched admin does not match created admin")
	}
}

func TestFetchAllAdmins(t *testing.T) {
	tx := testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	adminRepo := repository.Repository[models.Admin]{}
	adminRepo.SetDB(tx)

	authRepo := repository.AuthRepository{}
	authRepo.SetDB(tx)

	iterations := 5
	// Act
	for range iterations {
		testutils.CreateAdmin(t,tx,authRepo)
	}

	var admins []models.Admin
	if err := adminRepo.GetAll(&admins); err != nil {
		t.Errorf("Failed to fetch all admins: %v", err)
	}

	// Assert
	if len(admins) != iterations {
		t.Errorf("Expected at least %d admins, got %d", iterations, len(admins))
	}
}

func TestUpdateAdmin(t *testing.T) {
	tx := testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	adminRepo := repository.Repository[models.Admin]{}
	adminRepo.SetDB(tx)

	authRepo := repository.AuthRepository{}
	authRepo.SetDB(tx)

	// Act
	adminToCreate:=testutils.CreateAdmin(t,tx,authRepo)
	updatedAdminData := testutils.GenerateRandomAdmin()
	if err := adminRepo.Updates(updatedAdminData, adminToCreate.Id); err != nil {
		t.Errorf("Failed to update admin: %v", err)
	}

	// Assert
	var fetchedAdmin models.Admin
	if err := adminRepo.GetById(&fetchedAdmin, adminToCreate.Id); err != nil {
		t.Errorf("Failed to fetch updated admin: %v", err)
	}
	updatedAdminData.UserID = fetchedAdmin.UserID // UserID is not updated
	if fetchedAdmin.IsEmpty() {
		t.Errorf("Fetched admin is empty")
	}

	if !fetchedAdmin.IsEqual(updatedAdminData) {
		t.Errorf("Fetched admin does not match updated admin data")
	}
}
