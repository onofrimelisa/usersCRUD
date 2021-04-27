package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	Id int `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
}

func main() {
	r := gin.New()
	app := r.Group("api/v1/user")
	{
		app.GET("/",  GetUsers)
		app.GET("/:id",  GetUser)
		app.POST("/",  CreateUser)
	}

	err := r.Run(":8080")

	if err != nil {
		log.Fatalln(err)
	}
}

func GetUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")

	var user User
	db.First(&user, id)

	if user.Id != 0 {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func GetUsers(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var users []User
	db.Find(&users)

	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user User
	c.Bind(&user)

	if user.Firstname != "" && user.Lastname != "" {
		db.Create(&user)
		c.JSON(201, gin.H{"success": user})
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

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