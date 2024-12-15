package controllers

import (
	"awtopark/database"
	"awtopark/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateRF(c *fiber.Ctx) error {
	// RFID kart ID'sini almak için gelen veriyi parse et
	updateData := new(struct {
		RFIDCardID string `json:"rf_id_card_id"` // RFID kart ID'si
	})

	if err := c.BodyParser(updateData); err != nil {
		return c.Status(400).JSON("Veri formatı hatalı")
	}

	// RFID kartı bul
	rfidCard := models.RFIDCard{}
	if err := database.DB.Where("card_id = ?", updateData.RFIDCardID).First(&rfidCard).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON("RFID kart bulunamadı")
		}
		return c.Status(500).JSON("Sunucu hatası")
	}

	// Plakaya bağlı en son park kaydını bul
	park := models.ParkingEntry{}
	err := database.DB.Where("plate = ?", rfidCard.Plate).
		Order("id DESC").  // ID'yi büyükten küçüğe sıralar
		Limit(1).          // Yalnızca birinci kaydı alır
		First(&park).Error // İlk kaydı al
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON("Plakaya bağlı park kaydı bulunamadı")
		}
		return c.Status(500).JSON("Sunucu hatası")
	}

	// Çıkış zamanını şu anki zaman olarak ayarla
	exitTime := time.Now()
	formattedExitTime := exitTime.Format("2006-01-02 15:04:05")

	// Park kaydını güncelle (çıkış zamanını ayarla)
	if err := database.DB.Model(&park).Update("exit_time", formattedExitTime).Error; err != nil {
		return c.Status(500).JSON("Park kaydını güncellerken hata oluştu")
	}

	// İşlemi tamamla ve yanıt döndür
	return c.Status(200).JSON(fiber.Map{
		"message":  "Çıkış zamanı başarıyla güncellendi",
		"exitTime": exitTime.Format("2006-01-02 15:04:05"),
	})
}
