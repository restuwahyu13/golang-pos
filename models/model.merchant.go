package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelMerchant struct {
	ID         string `json:"id" gorm:"primary_key"`
	Name       string `json:"name" gorm:"type:varchar; not null"`
	Phone      string `json:"phone" gorm:"type:bigint; unique; not null"`
	Address    string `json:"address" gorm:"type:text; not null"`
	Logo       string `json:"logo" gorm:"type:varchar; not null"`
	Supplier   ModelSupplier
	SupplierID string    `json:"customer_id" gorm:"index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m *ModelMerchant) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.New().String()
	m.CreatedAt = time.Now().Local()
	return nil
}

func (m *ModelMerchant) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
