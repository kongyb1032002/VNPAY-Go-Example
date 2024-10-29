package factory

import (
	"vnpay-demo/src/internal/service"
	"vnpay-demo/src/pkg/auth"
	"vnpay-demo/src/pkg/hash"
)

type ServiceFactory struct {
	UserService service.UserService
	AuthService auth.Service
	HashService hash.Service
}

// NewServiceFactory sẽ tạo factory cho các service
func NewServiceFactory(repositoryFactory *RepositoryFactory, jwtSecret string) *ServiceFactory {
	hashService := hash.NewService()
	authService := auth.NewService(jwtSecret)

	return &ServiceFactory{
		UserService: service.NewUserService(repositoryFactory.UserRepository, hashService, authService),
		AuthService: authService,
		HashService: hashService,
	}
}
