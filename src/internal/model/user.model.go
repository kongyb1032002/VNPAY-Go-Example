package model

import (
	"log"
	"time"
	"vnpay-demo/src/pkg/hash"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	Entity
	Username             string    `json:"username" gorm:"unique;not null;" validate:"required"`
	FullName             string    `json:"full_name"`
	Address              string    `json:"address"`
	Country              string    `json:"country"`
	Province             string    `json:"province"`
	PostalCode           uint32    `json:"postal_code"`
	Password             string    `json:"password"`
	HashedPassword       string    `json:"hashed_password"`
	Confirmed            bool      `json:"confirmed" gorm:"default:0"`
	Enable2FA            bool      `json:"enable_2fa" gorm:"default:0"`
	PhoneNumber          string    `json:"phone_number" gorm:"not null" validate:"required,e164"`
	PhoneNumberConfirmed bool      `json:"phone_number_confirmed" gorm:"default:0"`
	Email                string    `json:"email" gorm:"unique;not null;" validate:"required,email"`
	EmailConfirmed       bool      `json:"email_confirmed" gorm:"default:0"`
	LockoutEnd           time.Time `json:"lockout_end" gorm:"default:null"`
	AccessFailedCount    uint8     `json:"access_failed_count" gorm:"default:0"`
	Roles                []Role    `json:"roles" gorm:"many2many:user_roles;"`
	RefreshToken         string    `json:"refreshToken"`
}

func (u *User) ValidateStruct() error {
	validate := validator.New()
	return validate.Struct(u)
}

func UserSeeder(db *gorm.DB) {
	// password1, _ := hashService.HashPassword("password123")
	hashService := hash.NewService()
	password, _ := hashService.HashPassword("password123")

	users := []User{
		{
			Username: "user1", FullName: "John Doe", Address: "123 Main St", Country: "USA", Province: "Texas", PostalCode: 75001, HashedPassword: password, Confirmed: true, Enable2FA: false, PhoneNumber: "+11234567890", PhoneNumberConfirmed: true, Email: "user1@example.com", EmailConfirmed: true,
			Roles: []Role{
				{Entity: Entity{ID: 1}},
			},
		},
		{
			Username: "user2", FullName: "Jane Smith", Address: "456 Elm St", Country: "USA", Province: "California", PostalCode: 90001, HashedPassword: password, Confirmed: true, Enable2FA: false, PhoneNumber: "+19876543210", PhoneNumberConfirmed: true, Email: "user2@example.com", EmailConfirmed: true,
			Roles: []Role{
				{Entity: Entity{ID: 1}},
			},
		},
		{
			Username: "user3", FullName: "Alice Johnson", Address: "789 Oak St", Country: "USA", Province: "New York", PostalCode: 10001, HashedPassword: password, Confirmed: true, Enable2FA: true, PhoneNumber: "+12345678901", PhoneNumberConfirmed: true, Email: "user3@example.com", EmailConfirmed: true,
			Roles: []Role{
				{Entity: Entity{ID: 1}},
			},
		},
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		log.Fatalf("fail to begin transaction: %s", err)
	}
	for _, user := range users {
		if err := db.FirstOrCreate(&user, User{Username: user.Username}).Error; err != nil {
			tx.Rollback()
			log.Fatalf("fail to seed user: %s", err)
		}
	}
	if err := tx.Commit().Error; err != nil {
		log.Fatalf("fail to commit transaction: %s", err)
	}
}

type UserStatus uint

const (
	Available         UserStatus = 1
	TemporarilyLocked UserStatus = 2
	PermanentlyLocked UserStatus = 3
	Unconfirmed       UserStatus = 4
	PendingApproval   UserStatus = 5
	Suspended         UserStatus = 6
	Deleted           UserStatus = 7
	Inactive          UserStatus = 8
)

var userStatusNames = []string{
	"Unknown",
	"Available",
	"Temporarily Locked",
	"Permanently Locked",
	"Unconfirmed",
	"Pending Approval",
	"Suspended",
	"Deleted",
	"Inactive",
}

func (status UserStatus) String() string {
	if status < Available || status > Inactive {
		return userStatusNames[0]
	}
	return userStatusNames[status]
}
