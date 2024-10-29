// router/auth_routes.go
package router

import (
	"vnpay-demo/src/internal/api"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, handler api.AuthHandler) {
	authRouter := r.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/sign-in", handler.SignIn).Methods("POST")
	authRouter.HandleFunc("/sign-up", handler.SignUp).Methods("POST")
	authRouter.HandleFunc("/change-password", handler.ChangePassword).Methods("POST")
	authRouter.HandleFunc("/forget-password", handler.ForgetPassword).Methods("POST")
	authRouter.HandleFunc("/profile", handler.UserProfile).Methods("GET")
}
