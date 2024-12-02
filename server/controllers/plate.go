package controllers

import (
	"awtopark/database"
	"awtopark/models"
	"encoding/json"
	"log"
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

func GetAllParkdata(c *fiber.Ctx) error {
	var parks []models.ParkingEntry
	if err := database.DB.Find(&parks).Error; err != nil {
		return c.Status(500).JSON("Server Interval Error")
	}
	return c.Status(200).JSON(fiber.Map{
		"data": parks,
	})
}

func GetAllParkdataWeb(c *websocket.Conn) {
	defer func() {
		c.Close()
	}()

	for {
		var parks []models.ParkingEntry
		if err := database.DB.Find(&parks).Error; err != nil {
			log.Println("Veritabanı hatası:", err)
			break
		}

		if err := c.WriteJSON(parks); err != nil {
			log.Println("WebSocket bağlantısı kapatıldı:", err)
			break
		}

		time.Sleep(5 * time.Second)
	}
}

func PostParkData(c *fiber.Ctx) error {
	data := new(models.ParkingEntry)
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON("Maglumat formaty yanlys")
	}
	data.EntryTime = time.Now().Format("2006-01-02 15:04:05")
	if err := database.DB.Create(&data).Error; err != nil {
		c.Status(500).JSON("Serveer INterval Error")
	}
	return c.Status(201).JSON(data)

}
func UpdateParkData(c *fiber.Ctx) error {
	updateData := new(struct {
		Plate string `json:"plate"`
	})
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(400).JSON("Maglumat formaty yanlys")
	}

	if updateData.Plate == "" {
		return c.Status(400).JSON("Plate boş geçilemez")
	}

	park := models.ParkingEntry{}
	err := database.DB.Where("plate = ?", updateData.Plate).First(&park).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON("Kayıt tapylmady")
		}
		return c.Status(500).JSON("Server Interval Error")
	}

	park.ExitTime = time.Now().Format("2006-01-02 15:04:05")

	layout := "2006-01-02 15:04:05"
	entryTime, err := time.Parse(layout, park.EntryTime)
	if err != nil {
		return c.Status(500).JSON("EntryTime formatı hatalı")
	}

	exitTime, err := time.Parse(layout, park.ExitTime)
	if err != nil {
		return c.Status(500).JSON("ExitTime formatı hatalı")
	}

	duration := exitTime.Sub(entryTime)
	minutes := int(duration.Minutes())

	var totalCost float64
	if minutes < 30 {
		totalCost = 0
	} else {
		hourlyRate := 1.0
		hours := float64(minutes) / 60.0
		calculatedCost := hours * hourlyRate
		totalCost = math.Round(calculatedCost*100) / 100
	}

	if err := database.DB.Model(&models.ParkingEntry{}).Where("plate = ?", updateData.Plate).Updates(map[string]interface{}{
		"exit_time": park.ExitTime,
	}).Error; err != nil {
		return c.Status(500).JSON("Server Interval Error")
	}

	payment := models.Payment{
		Plate: park.Plate,
		Date:  time.Now().Format("2006-01-02 15:04:05"),
		Money: totalCost,
	}
	if err := database.DB.Create(&payment).Error; err != nil {
		return c.Status(500).JSON("Server Interval Error")
	}

	return c.Status(200).JSON(fiber.Map{
		"plate":          park.Plate,
		"entryTime":      park.EntryTime,
		"exitTime":       park.ExitTime,
		"elapsedMinutes": minutes,
		"totalCost": func() interface{} {
			if totalCost == 0 {
				return "free"
			}
			return totalCost
		}(),
	})
}
func SearchParkData(c *fiber.Ctx) error {
	// Veritabanı bağlantısını doğrudan database.DB üzerinden al
	db := database.DB

	// Sorgu parametresini al
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Query parameter 'q' is required", // Query param 'q' is missing
		})
	}

	// Sonuçları filtrele
	var results []models.ParkingEntry
	// Perform a case-insensitive search using ILIKE for plate field
	if err := db.Where("plate ILIKE ?", "%"+query+"%").Find(&results).Error; err != nil {
		// Handle any errors that occur during the query
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch results", // General error when fetching data
		})
	}

	// Sonuçları döndür
	// Send the results back to the client
	return c.JSON(results) // Send array of ParkingEntry objects
}

func SendParkDataWebSocket(c *websocket.Conn) {
	for {
		var parks []models.ParkingEntry
		// WebSocket bağlantısı açık olduğu sürece verileri gönder
		if err := database.DB.Find(&parks).Error; err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(`{"error": "Server Interval Error"}`))
			return
		}

		// JSON formatına dönüştür
		response, _ := json.Marshal(fiber.Map{
			"data": parks,
		})

		// WebSocket üzerinden gönder
		if err := c.WriteMessage(websocket.TextMessage, response); err != nil {
			log.Println("Error sending message:", err)
			// WebSocket kapanmışsa, döngüden çık
			return
		}

		// Her 5 saniyede bir güncellenmiş veriyi gönder
		time.Sleep(5 * time.Second)
	}
}
