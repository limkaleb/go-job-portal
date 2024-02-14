package models

import "time"

type Talent struct {
	ID           uint   `gorm:"primaryKey" json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Email        string `gorm:"unique" json:"email,omitempty"`
	Password     []byte `json:"password,omitempty"`
	Applications []Application
	CreatedAt    time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time `gorm:"not null" json:"updated_at,omitempty"`
}
