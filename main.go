package main

import (
  "log"

  "github.com/gin-gonic/gin"
  "github.com/wooogi123/gin-mysql-boilerplage/controllers"
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
