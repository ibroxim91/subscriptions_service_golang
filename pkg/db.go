package pkg

import (
	"log"
	"subscriptions_service_golang/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.Author{}, &models.Book{}, &models.Comment{}); err != nil {
		log.Fatalf("migration error: %v", err)
	}
	return db
}
