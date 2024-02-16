package models

import (
	"time"
)

type Job struct {
	ID           uint          `gorm:"primaryKey" json:"id,omitempty"`
	Title        string        `json:"title,omitempty"`
	Description  string        `json:"description,omitempty"`
	Salary       uint          `json:"salary,omitempty"`
	Location     string        `json:"location,omitempty"`
	EmployerID   uint          `gorm:"index" json:"employer_id,omitempty"`
	Employer     *Employer     `json:"employer,omitempty"`
	Applications []Application `json:"applications,omitempty"`
	CreatedAt    time.Time     `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time     `gorm:"not null" json:"updated_at,omitempty"`
}

type PostJobRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Salary      string `json:"salary,omitempty"`
	Location    string `json:"location,omitempty"`
}

type PostJobResponse struct {
	ID          uint   `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Salary      uint   `json:"salary,omitempty"`
	Location    string `json:"location,omitempty"`
	EmployerID  uint   `json:"employer_id,omitempty"`
}
