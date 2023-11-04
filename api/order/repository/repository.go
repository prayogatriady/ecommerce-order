package repository

import (
	"context"

	"github.com/prayogatriady/ecommerce-module/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
}

type orderRepository struct {
	Db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		Db: db,
	}
}

func (r *orderRepository) Create(ctx context.Context, order *model.Order) error {
	return nil
}
