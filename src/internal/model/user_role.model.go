package model

type UserRole struct {
	Entity
	UserID uint64 `json:"user_id" gorm:"primaryKey;index:idx_user_role"`
	RoleID uint64 `json:"role_id" gorm:"primaryKey;index:idx_user_role"`
}
