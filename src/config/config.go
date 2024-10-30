package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpPort   string
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	JwtSecret  string
	DbPort     string
	OraclePwd  string
	OracleDb   string
	OracleHost string
	OraclePort string
	OracleUser string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	config := &Config{
		HttpPort:   os.Getenv("HTTP_PORT"),
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbName:     os.Getenv("DB_NAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		JwtSecret:  os.Getenv("JWT_SECRET"),
		DbPort:     os.Getenv("DB_PORT"),
		OraclePwd:  os.Getenv("ORACLE_PASSWORD"),
		OracleHost: os.Getenv("ORACLE_HOST"),
		OracleDb:   os.Getenv("ORACLE_DATABASE"),
		OracleUser: os.Getenv("ORACLE_USERNAME"),
		OraclePort: os.Getenv("ORACLE_PORT"),
	}
	return config, nil
}
