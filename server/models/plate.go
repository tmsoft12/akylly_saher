package models

type ParkingEntry struct {
	ID        int    `json:"id"`
	Plate     string `json:"plate"`
	EntryTime string `json:"entryTime"`
	ExitTime  string `json:"exitTime"`
}

type RFIDCard struct {
	ID      int     `json:"id"`      // Benzersiz ID (veritabanı için primary key)
	CardID  string  `json:"card_id"` // RFID kartın benzersiz kimliği
	Plate   string  `json:"plate"`   // Kartla ilişkilendirilmiş plaka
	Balance float64 `json:"balance"` // Kartın bakiyesi
}
