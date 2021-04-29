package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/onofrimelisa/usersCRUD/internal/models"
	"github.com/onofrimelisa/usersCRUD/internal/service"
	"net/http"
)

type UserController interface {
	CreateUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
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
*/
func (ctrl userController) UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.BaseUser

	err := ctrl.userService.GetUser(&user, id)

	if err != nil {
		_ = c.Error(err)
		return
	}

	var newUser models.BaseUser
	err = c.Bind(&newUser)

	if err != nil {
		_ = c.Error(err)
		return
	}

	err = ctrl.userService.UpdateUser(user, &newUser)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func (ctrl userController) GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.BaseUser

	err := ctrl.userService.GetUser(&user, id)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl userController) GetUsers(c *gin.Context) {
	var users []models.BaseUser

	err := ctrl.userService.GetUsers(&users)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, users)
}

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

	c.JSON(http.StatusCreated, user)
}

func (ctrl userController) DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.BaseUser

	err := ctrl.userService.GetUser(&user, id)

	if err != nil {
		_ = c.Error(err)
		return
	}

	err = ctrl.userService.DeleteUser(&user)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}
