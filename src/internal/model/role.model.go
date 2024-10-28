package model

import (
	"log"

	"gorm.io/gorm"
)

type Role struct {
	Entity
	Name        string `json:"name" gorm:"unique;not null" validate:"required"`
	Description string `json:"description"`
	Users       []User `gorm:"many2many:user_roles;"`
}

func RoleSeeder(db *gorm.DB) {
	roles := []Role{
		{Name: "Admin", Description: "System administrator with full access"},
		{Name: "Seller", Description: "Seller with rights to post and manage products"},
		{Name: "Buyer", Description: "Buyer with rights to purchase products and manage their orders"},
		{Name: "Moderator", Description: "Content moderator handling violations and content approval"},
		{Name: "Support", Description: "Customer support staff handling support requests"},
		{Name: "Guest", Description: "User not registered or logged in, can view products but cannot purchase"},
		{Name: "Delivery", Description: "Delivery personnel with access to orders for delivery"},
		{Name: "Finance", Description: "Finance staff with access to financial data and payment processing"},
	}

	tx := db.Begin()
	if err := tx.Error; err != nil {
		log.Fatalf("fail to begin transaction: %s", err)
	}
	for _, role := range roles {
		if err := tx.FirstOrCreate(&role, Role{Name: role.Name}).Error; err != nil {
			tx.Rollback()
			log.Fatalf("fail to seed roles: %s", err)
		}
	}
	if err := tx.Commit().Error; err != nil {
		log.Fatalf("fail to commit transaction: %s", err)
	}
}
