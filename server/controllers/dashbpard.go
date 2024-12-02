package controllers

import (
	"awtopark/database"
	"awtopark/models"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func CountEmptyExitTimes(c *fiber.Ctx) error {
	var emptyCount int64
	var totalCount int64
	var total sql.NullFloat64
	err := database.DB.Model(&models.Payment{}).Select("SUM(money)").Scan(&total).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Server Interval Error",
		})
	}
	totalPayments := 0.0
	if total.Valid {
		totalPayments = total.Float64
	}
	if err := database.DB.Model(&models.ParkingEntry{}).Where("exit_time = ? OR exit_time IS NULL", "").Count(&emptyCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not count entries with empty or NULL ExitTime",
		})
	}

	if err := database.DB.Model(&models.ParkingEntry{}).Count(&totalCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not count total parking entries",
		})
	}

	return c.JSON(fiber.Map{
		"empty_exit_count": emptyCount,
		"total_car_count":  totalCount,
		"totalPayments":    totalPayments,
	})
}

type DashboardData struct {
	EmptyExitCount int64   `json:"empty_exit_count"`
	TotalCarCount  int64   `json:"total_car_count"`
	TotalPayments  float64 `json:"totalPayments"`
}

func DashboardWebSocket(c *websocket.Conn) {
	for {
		var emptyCount int64
		var totalCount int64
		var total sql.NullFloat64
		err := database.DB.Model(&models.Payment{}).Select("SUM(money)").Scan(&total).Error
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(`{"error": "Server Interval Error"}`))
			return
		}

		totalPayments := 0.0
		if total.Valid {
			totalPayments = total.Float64
		}

		if err := database.DB.Model(&models.ParkingEntry{}).Where("exit_time = ? OR exit_time IS NULL", "").Count(&emptyCount).Error; err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(`{"error": "Could not count empty exit times"}`))
			return
		}

		if err := database.DB.Model(&models.ParkingEntry{}).Count(&totalCount).Error; err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(`{"error": "Could not count total parking entries"}`))
			return
		}
		data := DashboardData{
			EmptyExitCount: emptyCount,
			TotalCarCount:  totalCount,
			TotalPayments:  totalPayments,
		}
		response, _ := json.Marshal(data)
		c.WriteMessage(websocket.TextMessage, response)
		time.Sleep(5 * time.Second)
	}
}
