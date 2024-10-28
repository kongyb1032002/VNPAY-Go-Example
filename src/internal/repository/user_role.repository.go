package repository

import (
	"vnpay-demo/src/internal/model"

	"gorm.io/gorm"
)

type UserRoleRepository interface {
	Create(model *model.UserRole) error
	Update(model *model.UserRole) error
	Delete(ID uint64) error
	FindByUserID(ID uint64) (*[]model.UserRole, error)
	FindByRoleID(ID uint64) (*[]model.UserRole, error)
}

type userRoleRepository struct {
	db *gorm.DB
}

// Create implements UserRoleRepository.
func (r *userRoleRepository) Create(model *model.UserRole) error {
	return r.db.Create(&model).Error
}

// Delete implements UserRoleRepository.
func (r *userRoleRepository) Delete(ID uint64) error {
	panic("unimplemented")
}

// FindByRoleID implements UserRoleRepository.
func (r *userRoleRepository) FindByRoleID(ID uint64) (*[]model.UserRole, error) {
	var userRole []model.UserRole
	return &userRole, r.db.Where(&userRole, "role_id = ?", ID).Error
}

// FindByUserID implements UserRoleRepository.
func (r *userRoleRepository) FindByUserID(ID uint64) (*[]model.UserRole, error) {
	var userRole []model.UserRole
	return &userRole, r.db.Where(&userRole, "user_id = ?", ID).Error
}

// Update implements UserRoleRepository.
func (r *userRoleRepository) Update(model *model.UserRole) error {
	panic("unimplemented")
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleRepository{db}
}
