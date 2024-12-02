package controllers

import (
	"awtopark/database"
	"awtopark/models"
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateRF(c *fiber.Ctx) error {
	updateData := new(struct {
		RFIDCardID string `json:"rfid_card_id"` // RFID kart ID'si
	})

	// Gelen veriyi parse et
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(400).JSON("Veri formatı hatalı")
	}

	// RFIDCardID boş mu?
	if updateData.RFIDCardID == "" {
		return c.Status(400).JSON("RFID Kart ID gerekli")
	}

	// RFID kartı bul
	rfidCard := models.RFIDCard{}
	if err := database.DB.Where("card_id = ?", updateData.RFIDCardID).First(&rfidCard).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON("RFID kart bulunamadı")
		}
		return c.Status(500).JSON("Sunucu hatası")
	}

	// Kartın bir plakaya bağlı olup olmadığını kontrol et
	if rfidCard.Plate == "" {
		return c.Status(400).JSON("RFID kart bir plakaya bağlı değil")
	}

	// Plakaya göre park kaydını bul (çıkışı yapılmamış olanı buluyoruz)
	park := models.ParkingEntry{}
	err := database.DB.Where("plate = ? AND exit_time IS NULL", rfidCard.Plate).First(&park).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON("Çıkış yapılmamış park kaydı bulunamadı")
		}
		return c.Status(500).JSON("Sunucu hatası")
	}

	// Giriş zamanını al
	entryTime, err := time.Parse("2006-01-02 15:04:05", park.EntryTime)
	if err != nil {
		return c.Status(500).JSON("Giriş zamanı formatı hatalı")
	}

	// Çıkış zamanını belirle (şu anki zaman)
	exitTime := time.Now()

	// Geçen süreyi hesapla
	duration := exitTime.Sub(entryTime)
	minutes := int(duration.Minutes())

	// Ücreti hesapla
	var totalCost float64
	if minutes < 30 {
		totalCost = 0
	} else {
		hourlyRate := 1.0 // saatlik ücret
		hours := float64(minutes) / 60.0
		calculatedCost := hours * hourlyRate
		totalCost = math.Round(calculatedCost*100) / 100
	}

	// Bakiyeyi kontrol et
	if rfidCard.Balance < totalCost {
		return c.Status(400).JSON(fiber.Map{
			"error":          "Yeterli bakiye yok",
			"requiredCost":   totalCost,
			"currentBalance": rfidCard.Balance,
		})
	}

	// Bakiyeden düş
	rfidCard.Balance -= totalCost
	if err := database.DB.Save(&rfidCard).Error; err != nil {
		return c.Status(500).JSON("Bakiyeyi güncellerken hata oluştu")
	}

	// Park kaydını güncelle (çıkış zamanını güncelle)
	if err := database.DB.Model(&models.ParkingEntry{}).Where("id = ?", park.ID).Updates(map[string]interface{}{
		"exit_time": exitTime.Format("2006-01-02 15:04:05"),
	}).Error; err != nil {
		return c.Status(500).JSON("Park kaydını güncellerken hata oluştu")
	}

	// Ödeme kaydını oluştur
	payment := models.Payment{
		Plate:      rfidCard.Plate,
		Date:       time.Now().Format("2006-01-02 15:04:05"),
		Money:      totalCost,
		RFIDCardID: updateData.RFIDCardID,
	}
	if err := database.DB.Create(&payment).Error; err != nil {
		return c.Status(500).JSON("Ödeme kaydı oluşturulamadı")
	}

	// İşlemi tamamla ve yanıt döndür
	return c.Status(200).JSON(fiber.Map{
		"plate":          rfidCard.Plate,
		"entryTime":      park.EntryTime,
		"exitTime":       exitTime.Format("2006-01-02 15:04:05"),
		"elapsedMinutes": minutes,
		"totalCost": func() interface{} {
			if totalCost == 0 {
				return "free"
			}
			return totalCost
		}(),
		"remainingBalance": rfidCard.Balance,
	})
}
