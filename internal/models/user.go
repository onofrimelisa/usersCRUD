package models

import (
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type User interface {
	calculateAge()
}

type BaseUser struct {
	Id int `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
	BirthDate time.Time `gorm:"not null" form:"birthDate" json:"birthDate" time_format:"2006-01-02" time_utc:"1"`
}

func (u BaseUser) calculateAge() time.Duration{
	return time.Now().Sub(u.BirthDate)
}
