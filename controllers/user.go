package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"net/http"
	"github.com/lwnmengjing/goAdminBackand/forms"
	"github.com/lwnmengjing/goAdminBackand/models"
	"fmt"
)

type UserController struct {
	*BaseController
}

func (this *UserController) Init()  {

}

func (this *UserController) Create(c *gin.Context) {
	var userForm forms.User
	if err := c.BindJSON(&userForm); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": nil,
			"message": "params set failed",
		})
	}
	user := &models.User{
		Username: userForm.UserName,
	}
	if err := user.Insert(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"data": nil,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, user)
}

func (this *UserController) Get(c *gin.Context) {
	c.JSON(200, gin.H{"data":  "pong"})
}

func (this *UserController) Hello(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}

func (this *UserController) Register(c *gin.Context) {
	var register forms.Register
	if err := c.BindJSON(&register); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusUnprocessableEntity,
			"message": err.Error(),
			"data": nil,
		})
		return
	}

	user := models.User{
		Username: register.Username,
		RealName: register.RealName,
		Email: register.Email,
		Status: 1,
	}
	user.SetPassword(register.Password)
	if err := user.Insert(); err != nil {
		fmt.Println(err)

		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"message": err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusCreated,
		"message": nil,
		"data": user,
	})
}