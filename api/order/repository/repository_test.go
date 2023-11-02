package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/prayogatriady/ecommerce-order/api/model"
	"github.com/prayogatriady/ecommerce-order/database"
	"github.com/prayogatriady/ecommerce-order/utils/config"
	"github.com/prayogatriady/ecommerce-order/utils/constant"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	env *config.EnvVal
	Db  *gorm.DB
)

func TestMain(m *testing.M) {

	dir := fmt.Sprintf("../../../%s", constant.DIR_ENV)
	env = config.InitEnv(dir)
	db, _ := database.InitMysql(env)
	Db = db

	m.Run()
}

func TestConnection(t *testing.T) {
	assert.NotNil(t, Db)
}

func TestCreateRole(t *testing.T) {

	ctx := context.Background()

	role := []model.Role{
		{
			ID:          1,
			RoleName:    "Buyer",
			Description: "Pembeli",
			IsActive:    true,
		},
		{
			ID:          2,
			RoleName:    "Seller",
			Description: "Penjual",
			IsActive:    true,
		},
		{
			ID:          3,
			RoleName:    "Both",
			Description: "Penjual dan Pembeli",
			IsActive:    true,
		},
	}

	err := Db.WithContext(ctx).Save(&role).Error
	assert.Nil(t, err)

}

func TestCreateUser(t *testing.T) {

	ctx := context.Background()

	user := model.User{
		ID:       1,
		Username: "dobow",
		FullName: "Yogadobow",
		Password: "dobow",
		RoleId:   3,
		Balance:  100000,
		Email:    "dobow@gmail.com",
		Phone:    "085123456789",
		IsActive: true,
	}

	err := Db.WithContext(ctx).Save(&user).Error
	assert.Nil(t, err)

}

func TestCreateProduct(t *testing.T) {

	ctx := context.Background()

	product := []model.Product{
		{
			ID:          1,
			ProductName: "Purina ONE",
			Description: "Makanan Kucing Purina ONE",
			Quantity:    50,
			Price:       50000,
			UserId:      1,
			IsActive:    true,
		},
		{
			ID:          2,
			ProductName: "Royal Canin Kitten",
			Description: "Makanan Kucing Royal Canin",
			Quantity:    40,
			Price:       70000,
			UserId:      1,
			IsActive:    true,
		},
	}

	err := Db.WithContext(ctx).Save(&product).Error
	assert.Nil(t, err)

}

func TestCreateVoucher(t *testing.T) {

	ctx := context.Background()

	voucher := model.Voucher{
		ID:           1,
		Description:  "Diskon 11/11",
		MinPrice:     50000,
		QuantityAll:  100,
		QuantityUser: 1,
		ExpiredAt:    time.Now().Add(24 * time.Hour),
		IsActive:     true,
	}

	err := Db.WithContext(ctx).Save(&voucher).Error
	assert.Nil(t, err)

}

func TestCreateOrder(t *testing.T) {

	ctx := context.Background()

	order := model.Order{
		TotalPrice: 70000,
		UserId:     1,
		VoucherId:  1,
		OrderDetail: []model.OrderDetail{
			{
				OrderId:   1,
				ProductId: 2,
				Price:     70000,
				Quantity:  1,
			},
		},
	}

	err := Db.WithContext(ctx).Create(&order).Error
	assert.Nil(t, err)

}

func TestSelectOrder(t *testing.T) {

	ctx := context.Background()

	var order []model.Order
	err := Db.WithContext(ctx).Model(&model.Order{}).Preload("OrderDetail").Find(&order).Error
	assert.Nil(t, err)

}
