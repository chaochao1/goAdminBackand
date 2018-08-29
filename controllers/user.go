package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"net/http"
	"github.com/lwnmengjing/goAdminBackand/forms"
	"github.com/lwnmengjing/goAdminBackand/models"
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
		UserName: userForm.UserName,
		FirstName: userForm.FirstName,
		LastName: userForm.LastName,
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