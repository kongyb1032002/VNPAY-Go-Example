package middleware

import (
	"encoding/json"
	"net/http"
)

// ResponseWrapper là cấu trúc để gói dữ liệu phản hồi
type ResponseWrapper struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

// ResponseMiddleware để gói phản hồi
func ResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Tạo một responseWriter tùy chỉnh để ghi lại phản hồi
		writer := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(writer, r)

		// Ghi lại nội dung phản hồi từ writer
		var responseData interface{}
		if writer.responseData != nil {
			if err := json.Unmarshal(writer.responseData, &responseData); err != nil {
				responseData = nil // Nếu không giải mã được, đặt thành nil
			}
		}

		// Tạo dữ liệu phản hồi mới
		response := ResponseWrapper{
			Status: "OK",
			Code:   101,
			Data:   responseData,
		}

		// Ghi header và gửi phản hồi
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(writer.statusCode)
		json.NewEncoder(w).Encode(response)
	})
}

// responseWriter tùy chỉnh để ghi lại dữ liệu phản hồi
type responseWriter struct {
	http.ResponseWriter
	statusCode   int
	responseData []byte
}

// Ghi lại dữ liệu phản hồi
func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.responseData = append(rw.responseData, b...) // Ghi dữ liệu vào responseData
	return rw.ResponseWriter.Write(b)               // Ghi thẳng vào ResponseWriter
}

// Ghi mã trạng thái HTTP
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode                // Ghi mã trạng thái
	rw.ResponseWriter.WriteHeader(statusCode) // Ghi vào ResponseWriter
}
