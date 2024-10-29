package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vnpay-demo/src/internal/model"
	"vnpay-demo/src/internal/service"

	"github.com/gorilla/mux"
)

type userHandler struct {
	userService service.UserService
}

// Create implements CrudHandler.
func (u *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	panic("Unimplement method")
}

// Delete implements CrudHandler.
func (u *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := u.userService.Delete(id); err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "User deleted successfully"}
	json.NewEncoder(w).Encode(response)
}

// Detail implements CrudHandler.
func (u *userHandler) Detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := u.userService.GetByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// PageCount implements CrudHandler.
func (u *userHandler) PageCount(w http.ResponseWriter, r *http.Request) {
	count, err := u.userService.Total(nil)
	if err != nil {
		http.Error(w, "Error getting user count", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]uint64{"total_users": count})
}

// PagedList implements CrudHandler.
func (u *userHandler) PagedList(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	conditions := make(map[string]interface{})

	for key, values := range queryParams {
		if len(values) > 0 {
			value := values[0]
			switch key {
			case "offset", "limit":
				intValue, err := strconv.Atoi(value)
				if err == nil {
					conditions[key] = intValue
				}
			default:
				conditions[key] = value
			}
		}
	}

	users, err := u.userService.List(conditions)
	if err != nil {
		http.Error(w, "Error getting user list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Update implements CrudHandler.
func (u *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := u.userService.Update(&user); err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message": "User updated successfully",
		"user":    user,
	}
	json.NewEncoder(w).Encode(response)
}

func NewUserHandler(service service.UserService) CrudHandler {
	return &userHandler{userService: service}
}
