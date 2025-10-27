package testutils

import (
	"fmt"
	"log"
	"os"
	"github.com/houssybadr/lawyermanagement/backend/internal/database"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToTestDB() *gorm.DB {
	godotenv.Load("../.test.env")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to test database: " + err.Error())
	}
	return db
}

func SetUpTestDB() *gorm.DB {
	db := ConnectToTestDB()
	tx := db.Begin()
	database.Migrate(tx)
	return tx
}

func TearDownTestDB(tx *gorm.DB) {
	tx.Rollback()

	sqlDB, err := tx.DB()
	if err == nil {
		sqlDB.Close()
	}

}
