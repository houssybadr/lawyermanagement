package main

import (
	"test/internal/database"
	"test/internal/routers"
	"github.com/gin-gonic/gin"
)

func main(){
	database.ConnexionDB()
	database.Migrate(database.DB)
	db:=database.DB
	router:=gin.Default()
	routers.SetUpAuthRoutes(router,db)
	routers.SetUpAdminRoutes(router,db)
	routers.SetUpAvocatRouters(router,db)
	routers.SetUpClientRouters(router,db)
	routers.SetUpDossierRouters(router,db)
	routers.SetUpDocumentRouters(router,db)
	router.Run(":3000")
}