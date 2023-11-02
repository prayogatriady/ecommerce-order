package database

import (
	"fmt"
	"testing"

	"github.com/prayogatriady/ecommerce-order/utils/config"
	"github.com/prayogatriady/ecommerce-order/utils/constant"
	"github.com/stretchr/testify/assert"
)

var (
	env *config.EnvVal
)

func TestMain(m *testing.M) {

	dir := fmt.Sprintf("../%s", constant.DIR_ENV)
	env = config.InitEnv(dir)

	m.Run()
}

func Test_InitMysql(t *testing.T) {

	db, err := InitMysql(env)
	assert.NotNil(t, db)
	assert.Nil(t, err)

}
