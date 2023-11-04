package database

import (
	"log"
	"os"
	"testing"

	configM "github.com/prayogatriady/ecommerce-module/config"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	if err := configM.NewConfig(os.Getenv("APP_ENV"), "../"); err != nil {
		log.Fatal(err)
	}

	m.Run()
}

func Test_InitMysql(t *testing.T) {

	db, err := InitMysqlNew()
	assert.NotNil(t, db)
	assert.Nil(t, err)

}
