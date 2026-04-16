package models

type InvoiceDetail struct {
	ID        uint `gorm:"primaryKey"`
	InvoiceID uint
	ItemID    uint
	Quantity  int
	Price     int
	Subtotal  int
}