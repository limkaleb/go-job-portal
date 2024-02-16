package models

import (
	"time"
)

type Status string

const (
	Applied   Status = "applied"
	Interview Status = "interview"
	Accept    Status = "accept"
	Reject    Status = "reject"
)

type Application struct {
	ID             uint      `gorm:"primaryKey" json:"id,omitempty"`
	Status         Status    `gorm:"index" json:"status,omitempty"`
	SubmissionDate time.Time `gorm:"not null" json:"submission_date,omitempty"`
	TalentID       uint      `gorm:"uniqueIndex:job_talent_idx" json:"talent_id,omitempty"`
	JobID          uint      `gorm:"uniqueIndex:job_talent_idx" json:"job_id,omitempty"`
	Talent         *Talent 	`json:"talent,omitempty"`
	Job            *Job		`json:"job,omitempty"`
	CreatedAt      time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt      time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type ApplyJobResponse struct {
	ID             uint      `gorm:"primaryKey" json:"id,omitempty"`
	Status         Status    `gorm:"index" json:"status,omitempty"`
	SubmissionDate time.Time `gorm:"not null" json:"submission_date,omitempty"`
	TalentID       uint      `json:"talent_id,omitempty"`
	JobID          uint      `json:"job_id,omitempty"`
	CreatedAt      time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt      time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type UpdateApplicationRequest struct {
	Status         Status    `json:"status,omitempty"`
}
