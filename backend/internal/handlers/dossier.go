package handlers

import (
	"net/http"
	"strconv"
	"github.com/houssybadr/lawyermanagement/backend/internal/models"
	"github.com/houssybadr/lawyermanagement/backend/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DossierHandler struct {
	dossierRepository repository.Repository[models.Dossier]
}

func (d *DossierHandler) SetRepository(repo repository.Repository[models.Dossier]) {
	d.dossierRepository = repo
}

func (d *DossierHandler) GetAll(c *gin.Context) {
	var dossiers []models.Dossier
	if err := d.dossierRepository.GetAll(&dossiers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dossiers)
}

func (d *DossierHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var dossier models.Dossier
	if err := d.dossierRepository.GetById(&dossier, uint(id)); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Dossier not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dossier)
}

func (d *DossierHandler) Create(c *gin.Context) {
	var dossier models.Dossier
	if err := c.ShouldBindJSON(&dossier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := d.dossierRepository.Create(&dossier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dossier)
}

func (d *DossierHandler) GetByClientId(c *gin.Context) {
	client_id, err := strconv.Atoi(c.Param("client_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}
	var dossiers []models.Dossier
	if err := d.dossierRepository.GetByField(&dossiers, "client_id", uint(client_id)); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Dossiers not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dossiers)
}

func (d *DossierHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var dossier models.Dossier
	if err = c.ShouldBindJSON(&dossier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := d.dossierRepository.Updates(dossier, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var newDossier models.Dossier
	if err := d.dossierRepository.GetById(&newDossier, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newDossier)
}

func (d *DossierHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := d.dossierRepository.Delete(uint(id)); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Dossier not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dossier deleted successfully"})
}
