package routers

import (
	"github.com/houssybadr/lawyermanagement/backend/internal/handlers"
	"github.com/houssybadr/lawyermanagement/backend/internal/middlewares"
	"github.com/houssybadr/lawyermanagement/backend/internal/models"
	"github.com/houssybadr/lawyermanagement/backend/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func SetUpDossierRouters(r *gin.Engine,db *gorm.DB){
	dossierRepository := repository.Repository[models.Dossier]{}
	dossierRepository.SetDB(db)
	dossierHandler :=handlers.DossierHandler{}
	dossierHandler.SetRepository(dossierRepository)
	dossierRouters:=r.Group("/dossiers")
	dossierRouters.Use(middlewares.JwtAuthMiddleware())
	{
		dossierRouters.GET("/",dossierHandler.GetAll)
		dossierRouters.GET("/:id",dossierHandler.GetById)
		dossierRouters.GET("/client/:client_id",dossierHandler.GetByClientId)
		dossierRouters.POST("/",dossierHandler.Create)
		dossierRouters.PUT("/:id",dossierHandler.Update)
		dossierRouters.DELETE("/:id",dossierHandler.Delete)
	}
}