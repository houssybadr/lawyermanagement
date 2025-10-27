package routers

import (
	"test/internal/handlers"
	"test/internal/middlewares"
	"test/internal/models"
	"test/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func SetUpDocumentRouters(r *gin.Engine,db *gorm.DB){
	documerntRepository := repository.Repository[models.Document]{}
	documerntRepository.SetDB(db)
	documentHandler :=handlers.DocumentHandler{}
	documentHandler.SetRepository(documerntRepository)
	documerntRouters:=r.Group("/documents")
	documerntRouters.Use(middlewares.JwtAuthMiddleware())
	{
		documerntRouters.GET("/",documentHandler.GetAll)
		documerntRouters.GET("/:id",documentHandler.GetById)
		documerntRouters.GET("/file/:id",documentHandler.GetFile)
		documerntRouters.GET("/dossier/:dossier_id",documentHandler.GetByDossierId)
		documerntRouters.POST("/",documentHandler.Create)
		//documerntRouters.PUT("/:id",documentHandler.Update)
		documerntRouters.DELETE("/:id",documentHandler.Delete)
	}
}