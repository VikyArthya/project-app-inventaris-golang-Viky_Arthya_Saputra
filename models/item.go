package models

type Item struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Photo        string  `json:"photo"`
	Price        float64 `json:"price"`
	PurchaseDate string  `json:"purchase_date"`
	UsageDays    int     `json:"usage_days"`
	CategoryID   int     `json:"category_id"`
	CategoryName string  `json:"category_name"`
}
