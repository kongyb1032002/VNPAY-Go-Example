package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpPort   string
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
		return nil, err
	}
	config := &Config{
		HttpPort:   os.Getenv("HTTP_PORT"),
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbName:     os.Getenv("DB_NAME"),
		DbPassword: os.Getenv("DB_PASSWORd"),
	}
	return config, nil
}
