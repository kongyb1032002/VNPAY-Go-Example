package router

import (
	"vnpay-demo/src/internal/api"
	"vnpay-demo/src/middleware"

	"github.com/gorilla/mux"
)

// UserRoutes định nghĩa các route cho người dùng
func UserRoutes(r *mux.Router, handler api.CrudHandler) {
	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.Use(middleware.ResponseMiddleware)
	userRouter.HandleFunc("", handler.PagedList).Methods("GET").Headers("Content-Type", "application/json")
	userRouter.HandleFunc("", handler.Create).Methods("POST").Headers("Content-Type", "application/json")
	userRouter.HandleFunc("/{id}", handler.Update).Methods("PUT").Headers("Content-Type", "application/json")
	userRouter.HandleFunc("/{id}", handler.Detail).Methods("GET").Headers("Content-Type", "application/json")
	userRouter.HandleFunc("/{id}", handler.Delete).Methods("DELETE").Headers("Content-Type", "application/json")
}
