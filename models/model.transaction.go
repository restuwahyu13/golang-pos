package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ModelTransaction struct {
	ID           string `json:"id" gorm:"primary_key"`
	Customer     ModelCustomer
	CustomerID   string `json:"customer_id" gorm:"index"`
	Outlet       ModelOutlet
	OutletID     string         `json:"outlet_id" gorm:"index"`
	Products     pq.StringArray `json:"products" gorm:"type:text[]"`
	PurchaseDate time.Time      `json:"purchase_date"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func (m *ModelTransaction) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now()
	return nil
}

func (m *ModelTransaction) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now()
	return nil
}
