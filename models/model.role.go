package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type RoleAllowed string

const (
	admin    RoleAllowed = "admin"
	merchant RoleAllowed = "merchant"
	outlite  RoleAllowed = "outlet"
	supplier RoleAllowed = "supplier"
)

type ModelRole struct {
	ID         string         `json:"id" gorm:"primary_key"`
	RoleName   RoleAllowed    `json:"role_name" sql:"type:role_name"`
	RoleAccess pq.StringArray `json:"role_access" gorm:"type:text[]; not null"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

func (m *ModelRole) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.New().String()
	m.CreatedAt = time.Now().Local()
	return nil
}

func (m *ModelRole) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.New().String()
	m.UpdatedAt = time.Now().Local()
	return nil
}
