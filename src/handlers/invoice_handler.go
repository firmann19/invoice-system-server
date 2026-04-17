package handlers

import (
	"fmt"
	"time"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"fleetify-test/config"
	"fleetify-test/models"
)

func CreateInvoice(c *fiber.Ctx) error {
	var req CreateInvoiceRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	userID := uint(c.Locals("user_id").(int))

	err := config.DB.Transaction(func(tx *gorm.DB) error {

		invoice := models.Invoice{
			InvoiceNumber:   fmt.Sprintf("INV-%d", time.Now().Unix()),
			SenderName:      req.SenderName,
			SenderAddress:   req.SenderAddress,
			ReceiverName:    req.ReceiverName,
			ReceiverAddress: req.ReceiverAddress,
			TotalAmount:     0,
			CreatedBy:       userID,
		}

		if err := tx.Create(&invoice).Error; err != nil {
			return err
		}

		var total int = 0

		for _, itemReq := range req.Items {

			var item models.Item

			if err := tx.Where("code = ?", itemReq.Code).First(&item).Error; err != nil {
				return fmt.Errorf("item not found: %s", itemReq.Code)
			}

			subtotal := item.Price * itemReq.Quantity
			total += subtotal

			detail := models.InvoiceDetail{
				InvoiceID: invoice.ID,
				ItemID:    item.ID,
				Quantity:  itemReq.Quantity,
				Price:     item.Price,
				Subtotal:  subtotal,
			}

			if err := tx.Create(&detail).Error; err != nil {
				return err
			}
		}

		if err := tx.Model(&models.Invoice{}).
			Where("id = ?", invoice.ID).
			Update("total_amount", total).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Invoice created successfully",
	})
}