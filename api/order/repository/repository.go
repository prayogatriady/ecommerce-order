package repository

import (
	"context"

	modelM "github.com/prayogatriady/ecommerce-module/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	FindOrder(ctx context.Context, orderId int64) (*modelM.Order, error)
	FindOrders(ctx context.Context) (*[]modelM.Order, error)
	FindOrderDetail(ctx context.Context, orderDetailId int64, orderId int64) (*modelM.OrderDetail, error)
	FindOrderDetails(ctx context.Context, orderId int64) (*[]modelM.OrderDetail, error)
	CreateOrder(ctx context.Context, order *modelM.Order) (*modelM.Order, error)
	CreateOrderDetail(ctx context.Context, orderDetail *[]modelM.OrderDetail) (*[]modelM.OrderDetail, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) FindOrder(ctx context.Context, orderId int64) (*modelM.Order, error) {
	var order *modelM.Order
	err := r.db.WithContext(ctx).
		Model(&modelM.Order{}).
		Where("id = ?", orderId).
		Preload("OrderDetail").
		First(&order).Error
	return order, err
}

func (r *orderRepository) FindOrders(ctx context.Context) (*[]modelM.Order, error) {
	var orders *[]modelM.Order
	err := r.db.WithContext(ctx).Model(&modelM.Order{}).Preload("OrderDetail").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) FindOrderDetail(ctx context.Context, orderDetailId int64, orderId int64) (*modelM.OrderDetail, error) {
	var OrderDetail *modelM.OrderDetail
	err := r.db.WithContext(ctx).
		Model(&modelM.OrderDetail{}).
		Where("id = ? AND order_id = ?", orderDetailId, orderId).
		Preload("OrderDetail").
		First(&OrderDetail).Error
	return OrderDetail, err
}

func (r *orderRepository) FindOrderDetails(ctx context.Context, orderId int64) (*[]modelM.OrderDetail, error) {
	var OrderDetails *[]modelM.OrderDetail
	err := r.db.WithContext(ctx).
		Model(&modelM.OrderDetail{}).
		Where("order_id = ?", orderId).
		Preload("OrderDetail").
		First(&OrderDetails).Error
	return OrderDetails, err
}

func (r *orderRepository) CreateOrder(ctx context.Context, order *modelM.Order) (*modelM.Order, error) {
	err := r.db.WithContext(ctx).Create(&order).Error
	return order, err
}

func (r *orderRepository) CreateOrderDetail(ctx context.Context, orderDetail *[]modelM.OrderDetail) (*[]modelM.OrderDetail, error) {
	err := r.db.WithContext(ctx).Create(&orderDetail).Error
	return orderDetail, err
}
