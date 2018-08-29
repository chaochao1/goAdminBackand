package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/appleboy/gin-jwt"
	"log"
)

type BaseController struct {}

func (this *BaseController) Init()  {

}

func (this *BaseController) NotFound(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(http.StatusNotFound, gin.H{
		"code": http.StatusNotFound,
		"meesage": "Not Found!",
		"succes": 0,
	})
}