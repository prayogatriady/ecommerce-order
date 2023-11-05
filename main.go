package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	configM "github.com/prayogatriady/ecommerce-module/config"
	"github.com/prayogatriady/ecommerce-order/api/order/controller"
	redisRepository "github.com/prayogatriady/ecommerce-order/api/order/redis"
	"github.com/prayogatriady/ecommerce-order/api/order/repository"
	"github.com/prayogatriady/ecommerce-order/api/order/service"
	"github.com/prayogatriady/ecommerce-order/api/router"
	"github.com/prayogatriady/ecommerce-order/database"
	l "github.com/prayogatriady/ecommerce-order/utils/logger"
)

func main() {

	configM.NewConfig(os.Getenv("APP_ENV"), ".")

	l.InitLogger()

	db, err := database.InitMysqlNew()
	if err != nil {
		l.Logger.Fatal(err)
	}

	rdb := database.NewRedisClient()

	app, r := router.NewRouter()

	redisRepo := redisRepository.NewRedisRepository(rdb)
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(redisRepo, orderRepo)
	controller.NewOrderController(r, orderService)

	// Create an HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", configM.String("app.port", "8080")),
		Handler: app,
	}

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server failed to start: %v\n", err)
		}
	}()

	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Println("Server is shutting down...")

	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown:", err)
	} else {
		log.Println("Server is stopped gracefully 🔴")
	}
}
