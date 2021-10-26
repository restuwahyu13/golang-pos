package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelProduct struct {
	ID         string `json:"id" gorm:"primary_key"`
	Name       string `json:"name" gorm:"type:varchar; not null"`
	Image      string `json:"image" gorm:"type:varchar; not null"`
	SKU        uint64 `json:"sku" gorm:"type:bigint; not null; default=0"`
	Price      uint64 `json:"price" gorm:"type:bigint; not null; default=0"`
	Outlet     ModelOutlet
	OutletID   string `json:"outlet_id" gorm:"index"`
	Supplier   ModelSupplier
	SupplierID string    `json:"supplier_id" gorm:"index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m *ModelProduct) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now()
	return nil
}

func (m *ModelProduct) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now()
	return nil
}
