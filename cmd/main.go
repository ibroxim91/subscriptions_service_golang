package main

import (
	"log"
	"os"
	"subscriptions_service_golang/docs"
	"subscriptions_service_golang/internal/handlers"
	"subscriptions_service_golang/internal/middleware"
	"subscriptions_service_golang/internal/repositories"
	"subscriptions_service_golang/internal/services"
	"subscriptions_service_golang/pkg"
	"subscriptions_service_golang/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Subscription API
// @version 1.0
// @description API for managing subscriptions
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using system env")
	}

	logger.Init()
	defer logger.Log.Sync()

	dsn := os.Getenv("DB_DSN")
	database := pkg.Init(dsn) // db init
	repo := repositories.NewSubscriptionRepository(database)
	service := services.NewSubscriptionService(repo)
	handler := handlers.NewSubscriptionHandler(service)

	docs.SwaggerInfo.Title = "Subscription API"
	docs.SwaggerInfo.Description = "API for managing subscriptions"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authHandler := handlers.NewAuthHandler()
	r.POST("/login", authHandler.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(true))
	{
		auth.POST("/subscriptions", handler.Create)
		auth.PUT("/subscriptions/:id", handler.Update)
		auth.DELETE("/subscriptions/:id", handler.Delete)
	}

	// r.POST("/subscriptions", handler.Create)
	// r.PUT("/subscriptions/:id", handler.Update)
	// r.DELETE("/subscriptions/:id", handler.Delete)
	optional := r.Group("/")
	optional.Use(middleware.AuthMiddleware(false))
	{

		optional.GET("/subscriptions/:id", handler.GetByID)
		optional.GET("/subscriptions", handler.List)
		optional.GET("/subscriptions/total", handler.TotalPrice)
	}

	r.Run(":8080")
}
