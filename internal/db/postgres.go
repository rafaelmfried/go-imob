package db

import (
	fmt "fmt"
	config "imob/internal/config"
	entities "imob/internal/entities"
	log "log"

	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
)

func NewPostgres(cfg config.Config, logger interface{ Info(args ...interface{}) }) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Pass, cfg.Name, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil { log.Fatalf("db open: %v", err) }

	if err := db.AutoMigrate(&entities.User{}, &entities.Property{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}
	logger.Info("Connected to Postgres database")
	return db
}
