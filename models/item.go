package models

type Item struct {
	ID    uint   `gorm:"primaryKey"`
	Code  string `gorm:"unique"`
	Name  string
	Price int
}