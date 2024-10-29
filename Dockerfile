FROM golang:1.23.2 AS builder

WORKDIR /app

# Effectively tracks changes within your go.mod file
COPY go.mod go.sum ./

RUN go mod download

# Copies your source code into the app directory
COPY src/ ./src/

RUN go build -o /godocker ./src/main.go

EXPOSE 8080

CMD [ “/godocker” ] 