package router

import (
	"vnpay-demo/src/internal/factory" // Import HandlerFactory từ package factory

	"github.com/gorilla/mux"
)

// UseApiRouter sẽ nhận vào HandlerFactory thay vì từng handler
func UseApiRouter(handlerFactory *factory.HandlerFactory) *mux.Router {
	r := mux.NewRouter()

	UserRoutes(r, handlerFactory.UserHandler)
	AuthRoutes(r, handlerFactory.AuthHandler)

	return r
}
