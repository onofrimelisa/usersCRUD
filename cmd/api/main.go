package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/onofrimelisa/usersCRUD/internal/controller"
	"github.com/onofrimelisa/usersCRUD/internal/repository"
	"github.com/onofrimelisa/usersCRUD/internal/service"
	"log"
	"net/http"
)

func main() {
	r := gin.New()

	repo := repository.NewUserRepository()
	srv := service.NewUserService(repo)
	ctrl := controller.NewUserController(srv)

	app := r.Group("api/v1")
	{
		app.GET("/users", ctrl.GetUsers, ErrorHandle())
		app.GET("/users/:id", ctrl.GetUser, ErrorHandle())
		/*app.GET("/users/age", ctrl.CalculateAgeUser, ErrorHandle())*/
		app.POST("/users", ctrl.CreateUser, ErrorHandle())
		app.PUT("/users/:id", ctrl.UpdateUser, ErrorHandle())
		app.DELETE("/users/:id", ctrl.DeleteUser, ErrorHandle())
	}

	err := r.Run(":8080")

	if err != nil {
		log.Fatalln(err)
	}
}

func ErrorHandle() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Next()

		lastErr := c.Errors.Last()

		if lastErr != nil {
			if lastErr.IsType(gin.ErrorTypePrivate) || lastErr.IsType(gin.ErrorTypeBind) {
				c.JSON(http.StatusBadRequest, lastErr)
				return
			}

			c.JSON(http.StatusInternalServerError, lastErr)
			return
		}
	}
}

