package initializers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ishtiaqhimel/oms/auth-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	var err error
	maxWaitTime := time.Minute * 5
	retryInterval := time.Second * 20
	startTime := time.Now()

	for {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		if time.Since(startTime) > maxWaitTime {
			log.Fatalf("Could not connect to the database within %v: %v", maxWaitTime, err)
		}

		log.Printf("Database connection failed. Retrying in %v... Error: %v", retryInterval, err)
		time.Sleep(retryInterval)
	}

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to auto-migrate User model: %v", err)
	}
}
