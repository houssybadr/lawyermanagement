package handlers

import (
	"net/http"
	"strconv"
	"test/internal/models"
	"test/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AvocatHandler struct{
	avocatRepository repository.Repository[models.Avocat]
}

func (a *AvocatHandler)SetRepository(repo repository.Repository[models.Avocat]){
	a.avocatRepository=repo
}

func (a *AvocatHandler)GetAll(c *gin.Context){
	var avocats []models.Avocat
	if err:=a.avocatRepository.GetAll(&avocats);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	}
	c.JSON(http.StatusOK,avocats)
}

func (a *AvocatHandler)GetById(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID"})
		return
	}
	var avocat models.Avocat
	if err=a.avocatRepository.GetById(&avocat,uint(id));err!=nil{
		if err==gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound,gin.H{"error":"Avocat not found"})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,avocat)
}

func (a *AvocatHandler) GetByAdminId(c *gin.Context){
	admin_id,err:=strconv.Atoi(c.Param("admin_id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid admin ID"})
		return
	}
	var avocats []models.Avocat
	if err:=a.avocatRepository.GetByField(&avocats,"admin_id",uint(admin_id));err!=nil{
		if err==gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound,gin.H{"error":"Avocats not found"})
			return
		}	
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,avocats)
}

func (a *AvocatHandler) Update(c *gin.Context){
	var avocat models.Avocat
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID"})
		return
	}
	if err:=c.ShouldBindJSON(&avocat);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	if err:=a.avocatRepository.Updates(avocat,uint(id));err!=nil{
		if err==gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound,gin.H{"error":"Avocat not found"})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	if err:=a.avocatRepository.GetById(&avocat,uint(id));err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Avocat not found"})
		return
	}
	c.JSON(http.StatusOK,avocat)
}