package factory

import (
	"vnpay-demo/src/internal/repository"

	"gorm.io/gorm"
)

type RepositoryFactory struct {
	UserRepository repository.UserRepository
}

// NewRepositoryFactory sẽ tạo factory cho các repository
func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
	return &RepositoryFactory{
		UserRepository: repository.NewUserRepository(db),
	}
}
