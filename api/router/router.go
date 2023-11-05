package router

import "github.com/gin-gonic/gin"

func NewRouter() (*gin.Engine, *gin.RouterGroup) {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	return app, app.Group("/api")

}
