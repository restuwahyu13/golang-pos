package schemes

import "time"

type SchemeTransaction struct {
	ID           string    `json:"id" validate:"uuid"`
	CustomerID   string    `json:"customer_id" validate:"required,uuid"`
	OutletID     string    `json:"outlet_id" validate:"required,uuid"`
	Products     []string  `json:"products" validate:"required"`
	PurchaseDate time.Time `json:"purchase_date"`
}
