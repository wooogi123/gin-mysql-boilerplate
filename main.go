package main

import (
	"log"

	"./controllers"
	"github.com/gin-gonic/gin"
)

var userControl = new(controllers.UserController)

func main() {
	router := gin.Default()

	Users := router.Group("/Users")
	{
		Users.POST("/SignIn", userControl.SignIn)
		Users.POST("/SignUp", userControl.SignUp)
		Users.POST("/Update", userControl.Update)
	}

	log.Fatal(router.Run())
}
