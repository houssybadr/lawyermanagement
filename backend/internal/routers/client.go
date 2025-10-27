package routers

import (
	"test/internal/handlers"
	"test/internal/middlewares"
	"test/internal/models"
	"test/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func SetUpClientRouters(r *gin.Engine,db *gorm.DB){
	clientRepository := repository.Repository[models.Client]{}
	clientRepository.SetDB(db)
	clientHandler :=handlers.ClientHandler{}
	clientHandler.SetRepository(clientRepository)
	clientRouters:=r.Group("/clients")
	clientRouters.Use(middlewares.JwtAuthMiddleware())
	{
		clientRouters.GET("/",clientHandler.GetAll)
		clientRouters.GET("/:id",clientHandler.GetById)
		clientRouters.GET("/avocat/:avocat_id",clientHandler.GetByAvocatId)
		clientRouters.POST("/",clientHandler.Create)
		clientRouters.PUT("/:id",clientHandler.Update)
		clientRouters.DELETE("/:id",clientHandler.Delete)
	}
}