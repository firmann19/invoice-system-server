package database

import (
    "log"

    "fleetify-test/models"
    "gorm.io/gorm"
)

func SeedItems(db *gorm.DB) {
    items := []models.Item{
        {
            Code:  "BRG-001",
            Name:  "Barang Satu",
            Price: 10000,
        },
        {
            Code:  "BRG-002",
            Name:  "Barang Dua",
            Price: 20000,
        },
    }

    for _, item := range items {
        var existing models.Item

        err := db.Where("code = ?", item.Code).First(&existing).Error
        if err == nil {
            continue
        }

        if err := db.Create(&item).Error; err != nil {
            log.Println("❌ Failed seeding item:", err)
        } else {
            log.Println("✅ Seeded:", item.Code)
        }
    }
}