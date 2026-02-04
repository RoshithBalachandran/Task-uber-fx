package models

type User struct {
	ID    uint `gorm:"primary key"`
	Name  string
	Email string `gorm:"unique"`
}
