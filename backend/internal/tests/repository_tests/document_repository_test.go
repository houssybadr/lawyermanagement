package tests

import (
	"test/internal/models"
	"test/internal/repository"
	"test/internal/tests/testutils"
	"testing"
	//"fmt"
)

func TestCreateDocument(t *testing.T){
	tx :=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	documentRepo:=repository.Repository[models.Document]{}
	documentRepo.SetDB(tx)

	//Act 
	createdDocument:=testutils.CreateDocument(t,tx,documentRepo)

	//Assert
	var fetchedDocument models.Document
	if err:=documentRepo.GetById(&fetchedDocument,createdDocument.Id);err!=nil{
		t.Errorf("Failed to fetch document")
	}
	if createdDocument.IsEmpty(){
		t.Errorf("Feched document is empty")
	}
	if !fetchedDocument.IsEqual(createdDocument){
		t.Errorf("Fetched document does not match created document")
	}
}

func TestFetchDocumentById(t *testing.T){
	tx :=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	documentRepo:=repository.Repository[models.Document]{}
	documentRepo.SetDB(tx)

	//Act 
	createdDocument:=testutils.CreateDocument(t,tx,documentRepo)
	var fetchedDocument models.Document
	if err:=documentRepo.GetById(&fetchedDocument,createdDocument.Id);err!=nil{
		t.Errorf("Failed to fetch document")
	}
	//Assert
	
	if createdDocument.IsEmpty(){
		t.Errorf("Feched document is empty")
	}
	if !fetchedDocument.IsEqual(createdDocument){
		t.Errorf("Fetched document does not match created document")
	}
}

func TestFetchAllDocuments(t *testing.T){
	tx :=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	documentRepo:=repository.Repository[models.Document]{}
	documentRepo.SetDB(tx)

	iterations:=5

	//Act 
	for range iterations{
		testutils.CreateDocument(t,tx,documentRepo)
	}

	var fetchedDocuments []models.Document
	if err:=documentRepo.GetAll(&fetchedDocuments);err!=nil{
		t.Errorf("Failed to fetch all documents")
	}
	//Assert
	
	if len(fetchedDocuments)!=iterations{
		t.Errorf("Expected %d documents but got %document",iterations,len(fetchedDocuments))
	}
}

func TestDeleteDocument(t *testing.T){
	tx :=testutils.SetUpTestDB()
	defer testutils.TearDownTestDB(tx)

	documentRepo:=repository.Repository[models.Document]{}
	documentRepo.SetDB(tx)

	//Act 
	documentToDelete:=testutils.CreateDocument(t,tx,documentRepo)
	if err:=documentRepo.Delete(documentToDelete.Id);err!=nil{
		t.Errorf("Failed to delete document")
	}

	//
	var documentsCount int64
	documentRepo.Count(&documentsCount)
	if documentsCount!=0{
		t.Errorf("Failed to delete document")
	}
}