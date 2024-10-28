package router

import (
	"net/http"
	"strconv"
	"vnpay-demo/src/internal/api"
)

// UserRoutes định nghĩa các route cho người dùng
func TodoRoutes(handler api.TodoHandler) {
	// handler := api.NewTodoHandler()
	http.HandleFunc("/todo", func(rw http.ResponseWriter, r *http.Request) {
		// rw.Write([]byte("Todo Route"))
		rw.Header().Add("Content-Type", "application/json")
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		query := r.URL.Query()
		id, _ := strconv.Atoi(query.Get("id"))

		switch r.Method {
		case "GET":
			handler.GET(rw, r, id)
		case "POST":
			handler.POST(rw, r)
		case "PUT":
			handler.PUT(rw, r, id)
		case "DELETE":
			handler.DELETE(rw, r, id)
		}
	})
}
