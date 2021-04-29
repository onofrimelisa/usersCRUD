package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/onofrimelisa/usersCRUD/internal/models"
	"github.com/onofrimelisa/usersCRUD/internal/service"
	"net/http"
)

type UserController interface {
	CreateUser(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

/*func CalculateAgeUser(c *gin.Context) {
	// in progress
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.BaseUser
	db.First(&user, id)

	if user.Id != 0 {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
	fmt.Println()
}

func UpdateUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.BaseUser
	db.First(&user, id)

	if user.Id != 0 {
		var newUser models.BaseUser
		err := c.Bind(&newUser)

		if err != nil {
			log.Fatalln(err)
		}

		if !isEmpty(c, newUser.Firstname, newUser.Lastname) {
			result := models.BaseUser{
				Id:        user.Id,
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
	var user models.BaseUser
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

	var users []models.BaseUser
	db.Find(&users)

	c.JSON(200, users)
}*/

func (ctrl userController) CreateUser(c *gin.Context) {
	var user models.BaseUser
	err := c.Bind(&user)

	if err != nil {
		_ = c.Error(err)
		return
	}

	err = ctrl.userService.Create(&user)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": user})
}
/*
func DeleteUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.BaseUser
	db.First(&user, id)

	if user.Id != 0 {
		db.Delete(&user)
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

*/

