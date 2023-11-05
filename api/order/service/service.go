package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	configM "github.com/prayogatriady/ecommerce-module/config"
	modelM "github.com/prayogatriady/ecommerce-module/model"
	"github.com/prayogatriady/ecommerce-order/api/order/dto"
	redisRepo "github.com/prayogatriady/ecommerce-order/api/order/redis"
	"github.com/prayogatriady/ecommerce-order/api/order/repository"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

type OrderService interface {
	FindOrder(ctx context.Context, orderId int64) (*modelM.Order, error)
	CreateOrder(ctx context.Context, payload *dto.CreateUserDTO, dummy string) error
}

type orderService struct {
	redisRepo.RedisRepository
	repository.OrderRepository
}

func NewOrderService(redis redisRepo.RedisRepository, orderRepo repository.OrderRepository) OrderService {
	return &orderService{redis, orderRepo}
}

func (s *orderService) FindOrder(ctx context.Context, orderId int64) (*modelM.Order, error) {

	var order *modelM.Order

	redisKey := fmt.Sprintf("ORDER_%d", orderId)
	result, err := s.RedisRepository.Find(ctx, redisKey)
	if err != nil {
		return s.OrderRepository.FindOrder(ctx, orderId)
	}

	if err = json.Unmarshal([]byte(result), &order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderService) CreateOrder(ctx context.Context, payload *dto.CreateUserDTO, dummy string) (err error) {

	client := goredislib.NewClient(&goredislib.Options{
		Addr:     configM.String("redis.host", ""),
		Password: configM.String("redis.password", ""),
		DB:       0,
	})

	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	mutexname := fmt.Sprintf("VOUCHER_%d", payload.VoucherId)
	mutex := rs.NewMutex(mutexname, redsync.WithExpiry(30*time.Second), redsync.WithRetryDelay(40*time.Second))
	if err := mutex.Lock(); err != nil {
		return err
	}
	defer func() {
		client.Close()
	}()

	if dummy == "sleep" {
		time.Sleep(15 * time.Second)
	}

	order := &modelM.Order{
		TotalPrice: payload.TotalPrice,
		UserId:     payload.UserId,
		VoucherId:  payload.VoucherId,
	}

	order, err = s.OrderRepository.CreateOrder(ctx, order)
	if err != nil {
		return err
	}

	var orderDetails []modelM.OrderDetail
	for _, v := range payload.OrderDetail {
		orderDetail := modelM.OrderDetail{
			OrderId:   order.ID,
			ProductId: v.ProductId,
			Price:     v.Price,
			Quantity:  v.Quantity,
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	orderDetail, err := s.OrderRepository.CreateOrderDetail(ctx, &orderDetails)
	if err != nil {
		return err
	}
	order.OrderDetail = *orderDetail

	redisKey := fmt.Sprintf("ORDER_%d", order.ID)
	orderByte, _ := json.Marshal(order)
	if err = s.RedisRepository.Create(ctx, redisKey, orderByte); err != nil {
		return err
	}

	if ok, err := mutex.Unlock(); !ok || err != nil {
		return fmt.Errorf("unlock error: %v", err)
	}

	return nil
}
