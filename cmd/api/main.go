package main

import (
	"imob/internal/config"
	"imob/internal/db"
	"imob/internal/handlers"
	"imob/internal/logger"
	"imob/internal/repositories/gorm"
	"imob/internal/services"
	"log"

	gin "github.com/gin-gonic/gin"
)

func main() {
	config := config.Load()
	logger := logger.New(config)

	connection := db.NewPostgres(config, logger)

	userRepo := gorm.NewUserGorm(connection, logger)
	propertyRepo := gorm.NewPropertyGorm(connection, logger)

	userService := services.NewUserService(userRepo, logger)
	propertyService := services.NewPropertyService(propertyRepo, logger)

	userHandler := handlers.NewUserHandler(userService, logger)
	propertyHandler := handlers.NewPropertyHandler(propertyService, logger)

	runner := gin.Default()

	userHandler.Register(runner)
	propertyHandler.Register(runner)
	
	if err := runner.Run(config.HTTPAddr); err != nil {
		log.Fatalf("server error: %v", err)
	}

}