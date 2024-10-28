package main

import (
	"fmt"
	"log"
	"net/http"
	"vnpay-demo/src/config"
	"vnpay-demo/src/internal/api"
	"vnpay-demo/src/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	todoHandler := api.NewTodoHandler()
	router.TodoRoutes(todoHandler)

	fmt.Println("Server will run on port:", cfg.HttpPort)
	fmt.Println("Database host:", cfg.DbHost)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.HttpPort), nil))
}
