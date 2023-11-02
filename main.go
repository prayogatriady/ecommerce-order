package main

import (
	"github.com/prayogatriady/ecommerce-order/database"
	"github.com/prayogatriady/ecommerce-order/utils/config"
	l "github.com/prayogatriady/ecommerce-order/utils/logger"
)

var (
	env = config.Env
)

func main() {

	l.InitLogger()
	// l.Logger.Info("Starting")

	_, err := database.InitMysql(env)
	if err != nil {
		l.Logger.Fatal(err)
	}

	// app := gin.New()
	// app.Use(gin.Logger())
	// app.Use(gin.Recovery())
}
