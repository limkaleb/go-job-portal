package models

import "time"

type Talent struct {
	ID           uint          `gorm:"primaryKey" json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Email        string        `gorm:"unique" json:"email,omitempty"`
	Password     []byte        `json:"password,omitempty"`
	Applications []Application `json:"applications,omitempty"`
	CreatedAt    time.Time     `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time     `gorm:"not null" json:"updated_at,omitempty"`
}

type TalentRegisterRequest struct {
	Name     string `json:"name,omitempty"`
	Email    string `gorm:"unique" json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type TalentRegisterResponse struct {
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type TalentLoginRequest struct {
	Email    string `gorm:"unique" json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type GetTalentResponse struct {
	ID           uint          `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Email        string        `gorm:"unique" json:"email,omitempty"`
	Applications []Application `json:"applications,omitempty"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
}
