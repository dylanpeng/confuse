package router

import (
	"confuse/api/logic/control"
	apiMiddelware "confuse/api/logic/middleware"
	"confuse/common"
	ctrl "confuse/common/control"
	"confuse/common/middleware"
	"github.com/gin-gonic/gin"
)

var Router = &router{}

type router struct{}

func (r *router) GetIdentifier(ctx *gin.Context) string {
	return common.GetTraceId(ctx)
}

func (r *router) RegHttpHandler(app *gin.Engine) {
	app.Any("/health", ctrl.Health)
	app.Use(middleware.CheckEncoding)
	app.Use(middleware.CrossDomain)
	app.Use(middleware.Trace)

	apiGroup := app.Group("/api")
	{
		apiGroup.POST("/home/post", control.Home.HomePage)
		apiGroup.GET("/home/get", control.Home.HomePage)
	}

	userGroup := app.Group("/api/user")
	{
		userGroup.POST("/login", control.User.Login)
	}

	authUserGroup := app.Group("/api/user", apiMiddelware.Auth)
	{
		authUserGroup.POST("/info", control.User.GetInfo)
	}
}
