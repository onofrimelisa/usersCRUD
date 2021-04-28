package db

import (
	"github.com/jinzhu/gorm"
	. "github.com/onofrimelisa/usersCRUD/internal/user"
	"log"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./users_app.db")

	if err != nil {
		log.Fatalln(err)
	}

	db.LogMode(true)

	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	}

	return db
}
