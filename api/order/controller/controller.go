package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/ecommerce-order/api/order/dto"
	"github.com/prayogatriady/ecommerce-order/api/order/service"
)

type OrderController interface {
	CreateOrder(c *gin.Context)
	FindOrder(c *gin.Context)
}

type orderController struct {
	service.OrderService
}

func NewOrderController(r *gin.RouterGroup, s service.OrderService) {

	controller := &orderController{
		OrderService: s,
	}

	r = r.Group("/order")

	r.POST("/create", controller.CreateOrder)
	r.GET("/find", controller.FindOrder)
}

func (oc *orderController) CreateOrder(c *gin.Context) {

	var payload *dto.CreateUserDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := oc.OrderService.CreateOrder(context.Background(), payload, c.GetHeader("SLEEP")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "OK",
	})
}

func (oc *orderController) FindOrder(c *gin.Context) {

	query := c.Query("orderId")
	orderId, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := oc.OrderService.FindOrder(context.Background(), orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}
