package routers

import (
	"test/internal/handlers"
	"test/internal/middlewares"
	"test/internal/models"
	"test/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func SetUpAdminRoutes(r *gin.Engine,db *gorm.DB){
	adminRepository:=repository.Repository[models.Admin]{}
	adminRepository.SetDB(db)
	adminHandler:=handlers.AdminHandler{}
	adminHandler.SetRepository(adminRepository)
	adminRoutes :=r.Group("/admins")
	adminRoutes.Use(middlewares.JwtAuthMiddleware(),middlewares.CheckAdminMiddleware())
	{
		adminRoutes.GET("/",adminHandler.GetAll)
		adminRoutes.GET("/:id",adminHandler.GetById)
		adminRoutes.PUT("/:id",adminHandler.Update)
	}
}
