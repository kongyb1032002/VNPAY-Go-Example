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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connect mysql database: %v", err)
	}

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
