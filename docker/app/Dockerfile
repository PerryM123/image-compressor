FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache bash curl git imagemagick

# Use modules
ENV GO111MODULE=on

# CMD ["sh", "-c", "ls -la src"]
# CMD ["go", "run", "./src/main.go"]
CMD ["sh", "-c", "go run ./src/main.go"]