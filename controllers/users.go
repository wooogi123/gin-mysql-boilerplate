package controllers

import (
  "log"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/wooogi123/gin-mysql-boilerplate/forms"
  "github.com/wooogi123/gin-mysql-boilerplate/models"
)

type UserController struct{}

var userModel = new (models.User)

func (u UserController) SignIn(c *gin.Context) {
  model := new(forms.UserSignIn)
  if err := c.ShouldBindJSON(&model); err != nil {
    log.Println("Data not Bind!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": "Data not Bind!",
    })
    return
  }
  user, err := userModel.SignIn(*model)
  if err != nil {
    log.Println("Sign In Error!")
    c.JSON(http.StatusUnauthorized, gin.H{
      "status": "error",
      "content": err,
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
    "content": user,
  })
  return
}

func (u UserController) SignUp(c *gin.Context) {
  model := new(forms.UserSignUp)
  if err := c.ShouldBindJSON(&model); err != nil {
    log.Println("Data not Bind!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": "Data not Bind!",
    })
    return
  }
  user, err := userModel.SignUp(*model)
  if err != nil {
    log.Println("Sign Up Error!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": err,
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
    "content": user,
  })
  return
}

func (u UserController) Update(c *gin.Context) {
  model := new(forms.UserUpdate)
  if err := c.ShouldBindJSON(&model); err != nil {
    log.Println("Data not Bind!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": "Data not Bind!",
    })
    return
  }
  user, err := userModel.Update(*model)
  if err != nil {
    log.Println("Update Error!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": err,
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
    "content": user,
  })
  return
}
