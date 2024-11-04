package main

import (
	"fmt"
	"log"
	"net/http"
	"vnpay-demo/src/config"
	"vnpay-demo/src/internal/factory"
	"vnpay-demo/src/internal/model"
	"vnpay-demo/src/internal/router"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := config.LoadMysqlConfig(cfg)
	if err != nil {
		log.Fatalf("Error loading Mysql: %v", err)
	}

	// db, err := config.LoadOracleConfig(cfg)
	// if err != nil {
	// 	log.Fatalf("Error loading database: %v", err)
	// }

	// Tự động tạo các bảng
	err = db.AutoMigrate(&model.User{}, &model.Role{}) // Thêm các model khác nếu cần
	if err != nil {
		log.Fatalf("Error database migrate: %v", err)
	}
	repositoryFactory := factory.NewRepositoryFactory(db)
	serviceFactory := factory.NewServiceFactory(repositoryFactory, cfg.JwtSecret)
	handlerFactory := factory.NewHandlerFactory(serviceFactory)

	r := router.UseApiRouter(handlerFactory)

	fmt.Println("Server will run on port:", cfg.HttpPort)
	fmt.Println("Database host:", cfg.DbHost)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.HttpPort), r))
}
