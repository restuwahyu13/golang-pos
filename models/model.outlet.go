package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelOutlet struct {
	ID         string `json:"id" gorm:"primary_key"`
	Name       string `json:"name" gorm:"type:varchar; not null"`
	Phone      uint64 `json:"phone" gorm:"type:bigint; unique; not null"`
	Address    string `json:"address" gorm:"type:text; not null"`
	Merchant   ModelMerchant
	MerchantID string    `json:"merchant_id" gorm:"index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m *ModelOutlet) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now()
	return nil
}

func (m *ModelOutlet) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now()
	return nil
}
