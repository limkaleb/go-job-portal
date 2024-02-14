package models

type Talent struct {
	Id       uint
	Name     string
	Email    string `gorm:"unique"`
	Password []byte
}