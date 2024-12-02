package models

type Payment struct {
	ID         uint    `gorm:"primaryKey"`
	Plate      string  `json:"plate"`
	Date       string  `json:"date"`
	Money      float64 `json:"money"`
	RFIDCardID string  `json:"rfidcardid"`
}
