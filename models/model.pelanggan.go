package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelCustomer struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar; not null"`
	Phone     uint64    `json:"phone" gorm:"type:bigint; unique; not null"`
	Address   string    `json:"address" gorm:"type:text; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *ModelCustomer) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now()
	return nil
}

func (m *ModelCustomer) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now()
	return nil
}
