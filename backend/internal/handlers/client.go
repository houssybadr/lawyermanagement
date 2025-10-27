package handlers

import (
	"github.com/houssybadr/lawyermanagement/backend/internal/models"
	"github.com/houssybadr/lawyermanagement/backend/internal/repository"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClientHandler struct{
	clientRepository repository.Repository[models.Client]
}

func (cl *ClientHandler) SetRepository(repo repository.Repository[models.Client]){
	cl.clientRepository=repo
}

func (cl *ClientHandler)GetAll(c *gin.Context){
	var clients []models.Client
	if err:=cl.clientRepository.GetAll(&clients);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,clients)
}

func (cl *ClientHandler)GetById(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID"})
		return
	}
	var client models.Client
	if err=cl.clientRepository.GetById(&client,uint(id));err!=nil{
		if err==gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound,gin.H{"error":"Client not found"})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,client)
}

func (cl *ClientHandler) GetByAvocatId(c *gin.Context){
	avocat_id,err:=strconv.Atoi(c.Param("avocat_id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid avocat ID"})
		return
	}
	var clients []models.Client
	if err:=cl.clientRepository.GetByField(&clients,"avocat_id",uint(avocat_id));err!=nil{
		if err==gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound,gin.H{"error":"Clients not found"})
			return
		}	
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,clients)
}

func (cl *ClientHandler)Create(c *gin.Context){
	var client models.Client
	if err:=c.ShouldBindJSON(&client);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	if err:=cl.clientRepository.Create(&client);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	var newClient models.Client
	if err:=cl.clientRepository.GetById(&newClient,client.Id);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,newClient)
}

func (cl *ClientHandler)Update(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID"})
		return
	}
	var client models.Client
	if err=c.ShouldBindJSON(&client);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	if err:=cl.clientRepository.Updates(client,uint(id));err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	var newClient models.Client
	if err:=cl.clientRepository.GetById(&newClient,uint(id));err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,newClient)
}

func (cl *ClientHandler)Delete(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID"})
		return
	}
	if err:=cl.clientRepository.Delete(uint(id));err!=nil{
		if err==gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound,gin.H{"error":"Client not found"})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"Client deleted successfully"})
}
