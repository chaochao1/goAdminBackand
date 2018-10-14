package controllers

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type BaseController struct{}

func (this *BaseController) Init() {

}

func (this *BaseController) NotFound(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusNotFound,
		"meesage": "Not Found!",
		"succes":  0,
	})
}
