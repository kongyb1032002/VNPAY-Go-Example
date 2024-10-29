package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vnpay-demo/src/internal/model"
	"vnpay-demo/src/internal/request"
	"vnpay-demo/src/internal/service"
	"vnpay-demo/src/pkg/mapper"
)

type AuthHandler interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
	ForgetPassword(w http.ResponseWriter, r *http.Request)
	UserProfile(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	service service.UserService
}

// ChangePassword implements AuthHandler.
func (a *authHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var request request.ChangePassword

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Giả sử bạn có một ID người dùng được lưu trong token hoặc từ header Authorization
	userID := uint64(1) // Thay đổi này thành logic để lấy userID thực sự từ token hoặc session

	err := a.service.ChangePassword(userID, request.CurrentPassword, request.NewPassword)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Error changing password: %v", err)})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Change password successful!"}
	json.NewEncoder(w).Encode(response)
}

// ForgetPassword implements AuthHandler.
func (a *authHandler) ForgetPassword(w http.ResponseWriter, r *http.Request) {
	var request request.ForgetPassword

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	response := map[string]string{"message": fmt.Sprintf("Forget password successful for user: %v", request.Username)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SignIn implements AuthHandler.
func (a *authHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var request request.SignInRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	signInResponse, err := a.service.SignIn(request.Username, request.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Error signing in: %v", err)})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(signInResponse) // Trả về toàn bộ signInResponse
}

// SignUp implements AuthHandler.
func (a *authHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var request request.SignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	var user model.User

	if err := mapper.Map(&request, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Error mapping request: %v", err)})
		return
	}

	err := a.service.SignUp(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Error signing up: %v", err)})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message": "Sign up successful!",
		"user":    map[string]string{"username": user.Username, "email": user.Email},
	}
	json.NewEncoder(w).Encode(response)
}

// UserProfile implements AuthHandler.
func (a *authHandler) UserProfile(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
		return
	}

	// Phân tích token để lấy userID (giả sử có hàm DecodeToken)
	userID := uint64(1) // Thay đổi này thành logic để lấy userID thực sự từ token hoặc session

	user, err := a.service.GetByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("User not found: %v", err)})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{"user": user}
	json.NewEncoder(w).Encode(response)
}

func NewAuthHandler(service service.UserService) AuthHandler {
	return &authHandler{service: service}
}
