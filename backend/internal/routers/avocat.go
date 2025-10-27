package routers

import (
	"github.com/houssybadr/lawyermanagement/backend/internal/handlers"
	"github.com/houssybadr/lawyermanagement/backend/internal/middlewares"
	"github.com/houssybadr/lawyermanagement/backend/internal/models"
	"github.com/houssybadr/lawyermanagement/backend/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpAvocatRouters(r *gin.Engine,db *gorm.DB){
	avocatRepository:=repository.Repository[models.Avocat]{}
	avocatRepository.SetDB(db)
	avocatHandler :=handlers.AvocatHandler{}
	avocatHandler.SetRepository(avocatRepository)
	avocatRouters :=r.Group("/avocats")
	avocatRouters.Use(middlewares.JwtAuthMiddleware())
	{
		avocatRouters.GET("/",avocatHandler.GetAll)
		avocatRouters.GET(":id",avocatHandler.GetById)
		avocatRouters.GET("/admin/:admin_id",avocatHandler.GetByAdminId)
		avocatRouters.PUT("/:id",avocatHandler.Update)
	}
}