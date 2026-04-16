package services

import (
	"fleetify-test/models"

	"gorm.io/gorm"
)

func FindItemByCode(db *gorm.DB, code string) (*models.Item, error) {
	var item models.Item

	err := db.Where("code = ?", code).First(&item).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}