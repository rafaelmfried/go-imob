package config

import os "os"

type Config struct {
	Host string
	Port string
	Name   string
	User string
	Pass string
	HTTPAddr string
}

func Load() Config {
	return Config{
		Host:    getenv("DBHOST", "localhost"),
		Port:    getenv("DBPORT", "5432"),
		Name:    getenv("DBNAME", "postgres"),
		User:    getenv("DBUSER", "postgres"),
		Pass:   getenv("DBPASS", "password"),
		HTTPAddr: getenv("HTTP_ADDR", ":8080"),
	}
}

func getenv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}