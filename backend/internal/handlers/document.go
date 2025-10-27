package handlers

import (
	"io"
	"net/http"
	"strconv"
	"github.com/houssybadr/lawyermanagement/backend/internal/models"
	"github.com/houssybadr/lawyermanagement/backend/internal/repository"
	"time"
	"strings"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DocumentHandler struct{
	documentRepository repository.Repository[models.Document]
}

func (d *DocumentHandler)SetRepository(repo repository.Repository[models.Document]){
	d.documentRepository=repo
}

func (d *DocumentHandler) GetAll(c *gin.Context){
	var documents []models.Document
	if err:=d.documentRepository.GetAll(&documents);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,documents)
}

func (d *DocumentHandler) GetById(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID"})
		return
	}
	var document models.Document
	if err:=d.documentRepository.GetById(&document,uint(id));err!=nil{
		if err==gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound,gin.H{"error":"Document not found"})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,document)
}

func (a *DocumentHandler) GetByDossierId(c *gin.Context){
	dossier_id,err:=strconv.Atoi(c.Param("dossier_id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid dossier ID"})
		return
	}
	var documents []models.Document
	if err:=a.documentRepository.GetByField(&documents,"dossier_id",uint(dossier_id));err!=nil{
		if err==gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound,gin.H{"error":"Documents not found"})
			return
		}	
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,documents)
}

func (d *DocumentHandler) Create(c *gin.Context){
	file,header,err:=c.Request.FormFile("contenu")
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid file"})
		return
	}
	nom:=string(header.Filename)
	dateCreation:=time.Now()
	dossierId,err:=strconv.Atoi(c.PostForm("dossier_id"))
	var typeFichier models.TypeFichier
	nameSplits:=strings.Split(nom,".")
	extension:=strings.ToUpper(nameSplits[len(nameSplits)-1])
	if err:=typeFichier.FromString(extension);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	var typeDocumnet models.TypeDocument
	typeDocumnet.FromString(c.PostForm("type_document"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid dossier ID"})
		return
	}
	data,err:=io.ReadAll(file)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Cannot read file"})
		return
	}
	defer file.Close()

	document:=models.Document{
		Nom:nom,
		Contenu:data,
		DateCreation:dateCreation,
		DossierID:uint(dossierId),
		TypeFichier:typeFichier,
		TypeDocument:typeDocumnet,
	}
	if err:=d.documentRepository.Create(&document);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	var newDocument models.Document
	if err:=d.documentRepository.GetById(&newDocument,document.Id);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,newDocument)
}


func (d *DocumentHandler) GetFile(c *gin.Context){
	param:=c.Query("mode")
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID"})
		return
	}
	var document models.Document
	if err:=d.documentRepository.GetById(&document,uint(id));err!=nil{
		if err==gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound,gin.H{"error":"Document not found"})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	contentType,err:=document.TypeFichier.ToContentType()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Unknown file type"})
		return
	}
	
	if param=="download" || document.TypeFichier.ShouldBeDownloaded() {
		c.Header("Content-Disposition","attachment; filename="+document.Nom)
	}else{
		c.Header("Content-Disposition","inline; filename="+document.Nom)
	}
	c.Header("Content-Type",contentType)
	c.Data(http.StatusOK,"application/octet-stream",document.Contenu)
}

/*
func (d *DocumentHandler) Update(c *gin.Context){

}
*/
func (d *DocumentHandler) Delete(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID"})
		return
	}
	if err:=d.documentRepository.Delete(uint(id));err!=nil{
		if err==gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound,gin.H{"error":"Document not found"})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"Document deleted successfully"})
}