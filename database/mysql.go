package database

import (
	"fmt"

	"github.com/prayogatriady/ecommerce-order/utils/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(env *config.EnvVal) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.MySql.User, env.MySql.Password, env.MySql.Host, env.MySql.Port, env.MySql.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
