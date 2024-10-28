package repository

import (
	"vnpay-demo/src/internal/model"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *model.Role) error
	Update(role *model.Role) error
	Delete(roleID uint64) error
	FindByID(roleID uint64) (*model.Role, error)
	FindByName(name string) (*model.Role, error)
	List(conditions map[string]interface{}) (*[]model.Role, error)
	GetByIDs(ids []uint64) (*[]model.Role, error)
	HasUsersWithRole(roleID uint64) (bool, error)
	Total(conditions map[string]interface{}) (uint64, error)
}

type roleRepository struct {
	db *gorm.DB
}

func (r *roleRepository) HasUsersWithRole(roleID uint64) (bool, error) {
	var count int64
	err := r.db.Model(&model.UserRole{}).Where("role_id = ?", roleID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) List(conditions map[string]interface{}) (*[]model.Role, error) {
	roles := &[]model.Role{}
	db := r.db
	for key, value := range conditions {
		switch key {
		case "offset":
			db = db.Offset(value.(int))
		case "limit":
			db = db.Limit(value.(int))
		default:
			db = db.Where(key, value)
		}
	}
	err := db.Find(&roles).Error
	return roles, err
}

func (r *roleRepository) Create(role *model.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) Delete(roleID uint64) error {
	return r.db.Delete(&model.Role{}, roleID).Error
}

func (r *roleRepository) FindByID(roleID uint64) (*model.Role, error) {
	var role model.Role
	err := r.db.First(&role, "id = ?", roleID).Error
	return &role, err
}

func (r *roleRepository) FindByName(name string) (*model.Role, error) {
	var role model.Role
	err := r.db.First(&role, "name = ?", name).Error
	return &role, err
}

func (r *roleRepository) Update(role *model.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) GetByIDs(ids []uint64) (*[]model.Role, error) {
	roles := []model.Role{}
	db := r.db
	err := db.Where("id IN ?", ids).Find(&roles).Error
	return &roles, err
}

func (r *roleRepository) Total(conditions map[string]interface{}) (uint64, error) {
	var count int64
	db := r.db.Model(&model.Role{})
	for key, value := range conditions {
		db = db.Where(key, value)
	}
	err := db.Count(&count).Error
	return uint64(count), err
}
