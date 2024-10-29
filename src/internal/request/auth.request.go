package request

import "vnpay-demo/src/internal/model"

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type SignUpRequest struct {
	Username        string       `json:"username"`
	Password        string       `json:"password"`
	ConfirmPassword string       `json:"confirm_password"`
	FullName        string       `json:"full_name"`
	Address         string       `json:"address"`
	PhoneNumber     string       `json:"phone_number"`
	Email           string       `json:"email"`
	Roles           []model.Role `gorm:"many2many:user_roles;"`
}
type ChangePassword struct {
	ID              uint64 `json:"id"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}
type ForgetPassword struct {
	Username string `json:"username"`
}
