package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadMysqlConfig(cfg *Config) (*gorm.DB, error) {
	// Tạo DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort)
	// Kết nối đến MySQL mà không chỉ định database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Kiểm tra và tạo database nếu chưa tồn tại
	if err := db.Exec("CREATE DATABASE IF NOT EXISTS " + cfg.DbName).Error; err != nil {
		return nil, err
	}

	// Kết nối đến database đã tạo
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connect mysql database: %v", err)
	}

	return db, nil
}
