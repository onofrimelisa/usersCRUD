package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/onofrimelisa/usersCRUD/internal/controller"
	"github.com/onofrimelisa/usersCRUD/internal/repository"
	service2 "github.com/onofrimelisa/usersCRUD/internal/service"
	"log"
)

func main() {
	r := gin.New()

	repo := repository.NewUserRepository()
	srv := service2.NewUserService(repo)
	ctrl := controller.NewUserController(srv)

	app := r.Group("api/v1")
	{
		/*app.GET("/users", ctrl.GetUsers)
		app.GET("/users/:id", ctrl.GetUser)
		app.GET("/users/age", ctrl.CalculateAgeUser)*/
		app.POST("/users", ctrl.CreateUser)
		/*app.PUT("/users/:id", ctrl.UpdateUser)
		app.DELETE("/users/:id", ctrl.DeleteUser)*/
	}

	err := r.Run(":8080")

	if err != nil {
		log.Fatalln(err)
	}
}

