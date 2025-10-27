package handlers

import (
	"net/http"
	"strconv"
	"github.com/houssybadr/lawyermanagement/backend/internal/models"
	"github.com/houssybadr/lawyermanagement/backend/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminHandler struct {
	adminRepository repository.Repository[models.Admin]
}

func (a *AdminHandler) SetRepository(repo repository.Repository[models.Admin]) {
	a.adminRepository = repo
}


func (a *AdminHandler) GetAll(c *gin.Context) {
	var admins []models.Admin
	if err := a.adminRepository.GetAll(&admins); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admins)
}

func (a *AdminHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var admin models.Admin
	if err = a.adminRepository.GetById(&admin,uint(id)); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (a *AdminHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var admin models.Admin
	if err = c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = a.adminRepository.Updates(admin,uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = a.adminRepository.GetById(&admin,uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}




