package router

import (
	"net/http"
)

// UserRoutes định nghĩa các route cho người dùng
func UserRoutes() {
	http.HandleFunc("/users", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("User Route"))
	})
}
