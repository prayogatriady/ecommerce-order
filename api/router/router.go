package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/ecommerce-order/api/middleware"
)

func NewRouter() (*gin.Engine, *gin.RouterGroup) {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	return app, app.Group("/api", middleware.LoggerMiddleware)

}
