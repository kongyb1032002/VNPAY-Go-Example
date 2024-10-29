package factory

import (
	"vnpay-demo/src/internal/api"
)

type HandlerFactory struct {
	UserHandler api.CrudHandler
	AuthHandler api.AuthHandler
}

// NewHandlerFactory sẽ tạo factory cho các handler
func NewHandlerFactory(serviceFactory *ServiceFactory) *HandlerFactory {
	return &HandlerFactory{
		UserHandler: api.NewUserHandler(serviceFactory.UserService),
		AuthHandler: api.NewAuthHandler(serviceFactory.UserService),
	}
}
