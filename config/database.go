package config

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := "host=db user=postgres password=postgres dbname=fleetify port=5432 sslmode=disable"

	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Database connected")
			return
		}

		log.Println("⏳ Waiting database...", i+1)
		time.Sleep(2 * time.Second)
	}

	log.Fatal("❌ Failed to connect database:", err)
}