package config

import (
	"os"

	configM "github.com/prayogatriady/ecommerce-module/config"
)

func init() {

	configM.NewConfig(os.Getenv("APP_ENV"), ".")

}
