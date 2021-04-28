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
		app.PUT("/:id",  UpdateUser)
		app.DELETE("/:id",  DeleteUser)
	}

	err := r.Run(":8080")

	if err != nil {
		log.Fatalln(err)
	}
}

func UpdateUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user User
	db.First(&user, id)

	if user.Id != 0 {
		var newUser User
		c.Bind(&newUser)

		if !isEmpty(c, newUser.Firstname, newUser.Lastname) {
			result := User{
				Id: user.Id,
				Firstname: newUser.Firstname,
				Lastname:  newUser.Lastname,
			}

			db.Save(&result)
			c.JSON(200, gin.H{"success": result})
		}
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
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

	if !isEmpty(c, user.Firstname, user.Lastname) {
		db.Create(&user)
		c.JSON(201, gin.H{"success": user})
	}
}

func DeleteUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user User
	db.First(&user, id)

	if user.Id != 0 {
		db.Delete(&user)
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func isEmpty(c *gin.Context, fields ...string) bool {
	for _, field := range fields {
		if field == "" {
			c.JSON(422, gin.H{"error": "Fields are empty"})
			return true
		}
	}

	return false
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