package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	modelM "github.com/prayogatriady/ecommerce-module/model"
	mock_redis "github.com/prayogatriady/ecommerce-order/api/order/redis/mocks"
	mock_repository "github.com/prayogatriady/ecommerce-order/api/order/repository/mocks"
	l "github.com/prayogatriady/ecommerce-order/utils/logger"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_orderService_FindOrder(t *testing.T) {
	l.InitLogger()
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	type fields struct {
		RedisRepository func(*gomock.Controller) *mock_redis.MockRedisRepository
		OrderRepository func(*gomock.Controller) *mock_repository.MockOrderRepository
	}
	type args struct {
		ctx     context.Context
		orderId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *modelM.Order
		wantErr bool
	}{
		{
			name: "Test Success Redis Empty",
			fields: fields{
				RedisRepository: func(ctrl *gomock.Controller) *mock_redis.MockRedisRepository {
					mock := mock_redis.NewMockRedisRepository(ctrl)
					mock.EXPECT().Find(ctx, "ORDER_1").Return("", errors.New("Not Found"))
					return mock
				},
				OrderRepository: func(ctrl *gomock.Controller) *mock_repository.MockOrderRepository {
					mock := mock_repository.NewMockOrderRepository(ctrl)
					mock.EXPECT().FindOrder(ctx, int64(1)).Return(&modelM.Order{
						ID:          1,
						CreatedAt:   time.Time{},
						UpdatedAt:   time.Time{},
						DeletedAt:   gorm.DeletedAt{},
						OrderedAt:   time.Time{},
						TotalPrice:  1000,
						UserId:      123,
						VoucherId:   999,
						OrderDetail: nil,
						User:        nil,
					}, nil)
					return mock
				},
			},
			args: args{
				ctx:     ctx,
				orderId: 1,
			},
			want: &modelM.Order{
				ID:          1,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
				DeletedAt:   gorm.DeletedAt{},
				OrderedAt:   time.Time{},
				TotalPrice:  1000,
				UserId:      123,
				VoucherId:   999,
				OrderDetail: nil,
				User:        nil,
			},
			wantErr: false,
		},
		{
			name: "Test Success Redis Is Not Empty",
			fields: fields{
				RedisRepository: func(ctrl *gomock.Controller) *mock_redis.MockRedisRepository {
					mock := mock_redis.NewMockRedisRepository(ctrl)
					mock.EXPECT().Find(ctx, "ORDER_1").Return(`{"ID":1,"CreatedAt":"2024-01-01T12:34:56Z","UpdatedAt":"2024-01-01T12:34:56Z","DeletedAt":null,"OrderedAt":"2024-01-01T12:34:56Z","TotalPrice":100,"UserId":101,"VoucherId":201,"OrderDetail":null,"User":null}`, nil)
					return mock
				},
				OrderRepository: func(ctrl *gomock.Controller) *mock_repository.MockOrderRepository {
					mock := mock_repository.NewMockOrderRepository(ctrl)
					return mock
				},
			},
			args: args{
				ctx:     ctx,
				orderId: int64(1),
			},
			want: &modelM.Order{
				ID:          1,
				CreatedAt:   time.Date(2024, 1, 1, 12, 34, 56, 0, time.UTC),
				UpdatedAt:   time.Date(2024, 1, 1, 12, 34, 56, 0, time.UTC),
				DeletedAt:   gorm.DeletedAt{},
				OrderedAt:   time.Date(2024, 1, 1, 12, 34, 56, 0, time.UTC),
				TotalPrice:  100,
				UserId:      101,
				VoucherId:   201,
				OrderDetail: nil,
				User:        nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewOrderService(tt.fields.RedisRepository(ctrl), tt.fields.OrderRepository(ctrl))
			got, err := s.FindOrder(tt.args.ctx, tt.args.orderId)

			assert.Equal(t, err != nil, tt.wantErr)
			assert.EqualValues(t, got, tt.want)
		})
	}
}
