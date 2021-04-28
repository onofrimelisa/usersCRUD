package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onofrimelisa/usersCRUD/internal/db"
	. "github.com/onofrimelisa/usersCRUD/internal/user"
	"log"
)

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
		err := c.Bind(&newUser)

		if err != nil {
			log.Fatalln(err)
		}

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
	err := c.Bind(&user)

	if err != nil {
		log.Fatalln(err)
	}

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
