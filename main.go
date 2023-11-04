package main

import (
	"github.com/prayogatriady/ecommerce-order/database"
	l "github.com/prayogatriady/ecommerce-order/utils/logger"
)

func main() {

	l.InitLogger()

	_, err := database.InitMysqlNew()
	if err != nil {
		l.Logger.Fatal(err)
	}

	// app := gin.New()
	// app.Use(gin.Logger())
	// app.Use(gin.Recovery())
}
