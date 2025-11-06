package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/houssybadr/lawyermanagement/backend/internal/database"
	"github.com/houssybadr/lawyermanagement/backend/internal/routers"
)

func main(){
	database.ConnexionDB()
	//database.Migrate(database.DB)
	db:=database.DB
	router:=gin.Default()
	routers.SetUpAuthRoutes(router,db)
	routers.SetUpAdminRoutes(router,db)
	routers.SetUpAvocatRouters(router,db)
	routers.SetUpClientRouters(router,db)
	routers.SetUpDossierRouters(router,db)
	routers.SetUpDocumentRouters(router,db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}
	router.Run(":" + port)
}