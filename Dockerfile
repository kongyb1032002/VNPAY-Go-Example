# Bước 1: Sử dụng hình ảnh Go làm môi trường xây dựng
FROM golang:1.23.2 AS builder

# Thiết lập thư mục làm việc
WORKDIR /app

# Sao chép các file go.mod và go.sum
COPY go.mod go.sum ./

# Tải xuống các phụ thuộc
RUN go mod download

# Sao chép toàn bộ mã nguồn vào thư mục làm việc
COPY . .

# Sao chép file cấu hình .env.example vào thư mục làm việc
COPY .env.example .env

# Biên dịch ứng dụng và tạo file thực thi
RUN CGO_ENABLED=0 GOOS=linux go build -o /godocker ./src/main.go

# Bước 2: Tạo hình ảnh chính thức cho ứng dụng
FROM alpine:latest

# Sao chép file thực thi từ bước xây dựng
COPY --from=builder /godocker /godocker
# Sao chép file .env vào đường dẫn gốc
COPY --from=builder /app/.env /

# Thiết lập quyền thực thi cho file
RUN chmod +x /godocker

# Mở cổng 8080
EXPOSE 8080

# Chạy ứng dụng
CMD ["/godocker"]
