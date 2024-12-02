package models

type Stats struct {
	DailyAccount string `json:"daily_account"`
	TotalIncome  string `json:"total_income"`
	Vehicles     string `json:"vehicles"`
	FreeSpots    string `json:"free_spots"`
}
