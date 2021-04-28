package user

import (
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id int `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
}
