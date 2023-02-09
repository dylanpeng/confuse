package router

import (
	"confuse/api/logic/control"
	ctrl "confuse/common/control"
	"confuse/common/middleware"
	"github.com/gin-gonic/gin"
)

var Router = &router{}

type router struct{}

func (r *router) GetIdentifier(ctx *gin.Context) string {
	return "unknown"
}

func (r *router) RegHttpHandler(app *gin.Engine) {
	app.Any("/health", ctrl.Health)
	app.Use(middleware.CheckEncoding)
	app.Use(middleware.CrossDomain)

	apiGroup := app.Group("/api")
	{
		apiGroup.POST("/home/post", control.Home.HomePage)
		apiGroup.GET("/home/get", control.Home.HomePage)
	}

}
