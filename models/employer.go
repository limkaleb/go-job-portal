package models

import (
	"time"
)

type Employer struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `gorm:"unique" json:"email,omitempty"`
	Password  []byte    `json:"password,omitempty"`
	Jobs      []Job     `json:"jobs,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type EmployerRegisterRequest struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type EmployerRegisterResponse struct {
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type EmployerLoginRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type GetEmployerResponse struct {
	ID    uint   `gorm:"primaryKey" json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `gorm:"unique" json:"email,omitempty"`
	Jobs  []Job  `json:"jobs,omitempty"`
}
