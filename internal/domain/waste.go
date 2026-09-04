package domain

import "time"

type Waste struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	StockAvailability float64   `json:"stock_availability"`
	Costs             float64   `json:"costs"`
	Category          string    `json:"category"`
	CreatedAt         time.Time `json:"created_at"`
}
