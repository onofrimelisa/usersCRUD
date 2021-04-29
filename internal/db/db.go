package db

import (
	"github.com/jinzhu/gorm"
	. "github.com/onofrimelisa/usersCRUD/internal/models"
	"log"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./users_app.db")

	if err != nil {
		log.Fatalln(err)
	}

	db.LogMode(true)

	if !db.HasTable(&BaseUser{}) {
		db.CreateTable(&BaseUser{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&BaseUser{})
	}

	return db
}
