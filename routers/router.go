package routers

import (
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
	"github.com/lwnmengjing/goAdminBackand/controllers"
	"github.com/lwnmengjing/goAdminBackand/middlewares"
)

var (
	Router = gin.Default()
)

func init() {

	base := new(controllers.BaseController)
	user := new(controllers.UserController)
	authMiddleware := middlewares.AuthMiddleware

	Router.POST("/login", authMiddleware.LoginHandler)
	Router.POST("/register", user.Register)
	auth := Router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", user.Hello)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	}

	v1 := Router.Group("/v1", authMiddleware.MiddlewareFunc(), authz.NewAuthorizer(middlewares.Enforce))
	{
		v1.GET("/ping", user.Get)
		v1.POST("/user", user.Create)
	}

	Router.NoRoute(authMiddleware.MiddlewareFunc(), base.NotFound)
}
