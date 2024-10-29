package response

import (
	"time"
	"vnpay-demo/src/internal/model"
)

type User struct {
	model.Entity
	Username             string       `json:"username" gorm:"unique;not null;" validate:"required"`
	FullName             string       `json:"full_name"`
	Address              string       `json:"address"`
	Country              string       `json:"country"`
	Province             string       `json:"province"`
	PostalCode           uint32       `json:"postal_code"`
	Confirmed            bool         `json:"confirmed" gorm:"default:0"`
	Enable2FA            bool         `json:"enable_2fa" gorm:"default:0"`
	PhoneNumber          string       `json:"phone_number" gorm:"not null" validate:"required,e164"`
	PhoneNumberConfirmed bool         `json:"phone_number_confirmed" gorm:"default:0"`
	Email                string       `json:"email" gorm:"unique;not null;" validate:"required,email"`
	EmailConfirmed       bool         `json:"email_confirmed" gorm:"default:0"`
	LockoutEnd           time.Time    `json:"lockout_end"`
	AccessFailedCount    uint8        `json:"access_failed_count" gorm:"default:0"`
	Roles                []model.Role `json:"roles" gorm:"many2many:user_roles;"`
	StatusDetail         string       `json:"status_detail"`
}

type SignInResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
