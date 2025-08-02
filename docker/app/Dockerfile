FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache bash curl git imagemagick

RUN go install github.com/air-verse/air@latest

# Use modules
ENV GO111MODULE=on

CMD ["sh", "-c", "go run ./main.go"]