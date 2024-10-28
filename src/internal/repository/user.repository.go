package repository

import (
	"vnpay-demo/src/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(userID uint64) error
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByID(userID uint64) (*model.User, error)
	List(conditions map[string]interface{}) (*[]model.User, error)
	Total(conditions map[string]interface{}) (uint64, error)
	GetByIDs(ids []uint64) (*[]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// Total implements UserRepository.
func (r *userRepository) Total(conditions map[string]interface{}) (uint64, error) {
	var count int64
	db := r.db.Model(&model.User{})
	for key, value := range conditions {
		db = db.Where(key, value)
	}
	err := db.Count(&count).Error
	return uint64(count), err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) List(conditions map[string]interface{}) (*[]model.User, error) {
	user := &[]model.User{}
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

	if err := db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByID(userID uint64) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *model.User) error {
	if err := user.ValidateStruct(); err != nil {
		return err
	}
	return r.db.Create(user).Error
}

func (r *userRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.db.Preload("Roles").Where("username = ?", username).First(&user)
	return &user, result.Error
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (r *userRepository) Update(user *model.User) error {
	if err := user.ValidateStruct(); err != nil {
		return err
	}
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(userID uint64) error {
	return r.db.Delete(&model.User{}, userID).Error
}

func (r *userRepository) GetByIDs(ids []uint64) (*[]model.User, error) {
	users := []model.User{}
	err := r.db.Where("id IN ?", ids).Find(&users).Error
	return &users, err
}
