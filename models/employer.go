package models

type Employer struct {
	Id       uint
	Name     string
	Email    string `gorm:"unique"`
	Password []byte
}

