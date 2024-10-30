package main

import (
	"fmt"
	"log"
	"net/http"
	"vnpay-demo/src/config"
	"vnpay-demo/src/internal/factory"
	"vnpay-demo/src/internal/model"
	"vnpay-demo/src/internal/router"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Tạo DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort)
	fmt.Println(dsn)

	// Kết nối đến MySQL mà không chỉ định database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connect mysql database: %v", err)
	}

	// Kiểm tra và tạo database nếu chưa tồn tại
	if err := db.Exec("CREATE DATABASE IF NOT EXISTS " + cfg.DbName).Error; err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	// Kết nối đến database đã tạo
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connect mysql database: %v", err)
	}

	// Tự động tạo các bảng
	err = db.AutoMigrate(&model.User{}, &model.Role{}) // Thêm các model khác nếu cần
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	repositoryFactory := factory.NewRepositoryFactory(db)
	serviceFactory := factory.NewServiceFactory(repositoryFactory, cfg.JwtSecret)
	handlerFactory := factory.NewHandlerFactory(serviceFactory)

	r := router.UseApiRouter(handlerFactory)

	fmt.Println("Server will run on port:", cfg.HttpPort)
	fmt.Println("Database host:", cfg.DbHost)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.HttpPort), r))
}
