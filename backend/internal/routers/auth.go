package routers

import (
	"github.com/houssybadr/lawyermanagement/backend/internal/handlers"
	"github.com/houssybadr/lawyermanagement/backend/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpAuthRoutes(r *gin.Engine, db *gorm.DB) {
	authRepository := repository.AuthRepository{}
	authRepository.SetDB(db)
	authHandler := handlers.AuthHandler{}
	authHandler.SetRepository(authRepository)
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup/admin", authHandler.SignUpAdmin)
		authRoutes.POST("/signup/avocat", authHandler.SignUpAvocat)
		authRoutes.POST("/signin", authHandler.SignIn)
		authRoutes.POST("/change-password/:id", authHandler.ChangePassword)
		authRoutes.DELETE("/:id",authHandler.DeleteUser)
	}
}
