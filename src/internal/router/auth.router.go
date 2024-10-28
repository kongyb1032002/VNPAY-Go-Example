package router

import "net/http"

func AuthRoutes() {
	http.HandleFunc("/auth", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Auth Route"))
	})
}
