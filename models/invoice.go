package models

import "time"

type Invoice struct {
	ID              uint `gorm:"primaryKey"`
	InvoiceNumber   string
	SenderName      string
	SenderAddress   string
	ReceiverName    string
	ReceiverAddress string
	TotalAmount     int
	CreatedBy       uint
	CreatedAt       time.Time
}