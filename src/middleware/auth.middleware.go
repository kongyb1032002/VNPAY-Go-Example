package middleware

import (
	"net/http"
)

// AuthMiddleware kiểm tra xác thực
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// Kiểm tra xác thực ở đây
		// Nếu xác thực không hợp lệ
		// http.Error(rw, "Unauthorized", http.StatusUnauthorized)

		// Nếu xác thực hợp lệ
		next.ServeHTTP(rw, r)
	})
}
