FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache bash curl git imagemagick

# live reloading
RUN go install github.com/air-verse/air@latest

COPY app/go.mod ./

RUN go mod download

COPY app/ .

RUN go build -o main

CMD ["air"]