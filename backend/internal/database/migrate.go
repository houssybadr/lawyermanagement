package database
import (
	"gorm.io/gorm"
	"github.com/houssybadr/lawyermanagement/backend/internal/models"
	"log"
	"fmt"
)

func Migrate(db *gorm.DB){
	err:=db.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Avocat{},
		&models.Client{},
		&models.Client{},
		&models.Dossier{},
		&models.Document{},
	)
	if err!=nil{
		log.Fatal("Error during migration: ",err)
	}
	fmt.Println("Database migrated successfully")
}