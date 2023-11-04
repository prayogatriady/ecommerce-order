package database

import (
	"fmt"

	configM "github.com/prayogatriady/ecommerce-module/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysqlNew() (*gorm.DB, error) {

	var (
		user     = configM.String("mysql.user", "")
		password = configM.String("mysql.password", "")
		host     = configM.String("mysql.host", "")
		port     = configM.Int("mysql.port", 3306)
		name     = configM.String("mysql.name", "")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
