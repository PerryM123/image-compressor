package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

const maxFileSizeInMegabytes int64 = 10
const maxFileSizeForUploadImage int64 = maxFileSizeInMegabytes * 1024 * 1024 // 10 MB
var bearerTokenFromEnvFile = ""
// TODO: utilsディレクトリに移動した方がいいか検討中
type CompressionQuality int
const (
    LosslessCompression CompressionQuality = 100
	VeryHighCompression CompressionQuality = 95
	WebOptimizedCompression CompressionQuality = 85
	ModerateCompression CompressionQuality = 75
	SignificantCompression CompressionQuality = 60
	ExtremeCompression CompressionQuality = 10
)
type CompressResponse struct {
    Message        string `json:"message"`
    FileName       string `json:"fileName"`
    Base64Image    string `json:"base64Image"`
    OriginalSize   int64  `json:"originalSize"`
    CompressedSize int64  `json:"compressedSize"`
}
type HealthCheckResponse struct {
    Alive        bool `json:"alive"`
}
func handleErrorResponse(w http.ResponseWriter, response *CompressResponse, message string, statusCode int) {
    response.Message = message
    log.Println(message)
    w.WriteHeader(statusCode)
}
func compressImageHandler(w http.ResponseWriter, r *http.Request) {
    response := CompressResponse{}
    w.Header().Set("Content-Type", "application/json")
    defer func() {
        json.NewEncoder(w).Encode(response)
    }()
    authHeader := r.Header.Get("Authorization")
    tokenFromRequest := strings.TrimPrefix(authHeader, "Bearer ")
    if authHeader == "" {
		handleErrorResponse(w, &response, "Missing Authorization header", http.StatusUnauthorized)
        return
    }
    if !strings.HasPrefix(authHeader, "Bearer ") {
		handleErrorResponse(w, &response, "Invalid Authorization header format", http.StatusBadRequest)
        return
    }
    if tokenFromRequest != bearerTokenFromEnvFile {
		handleErrorResponse(w, &response, "Invalid Token Authorization", http.StatusUnauthorized)
        return
    }
	if r.Method != http.MethodPost {
		handleErrorResponse(w, &response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	file, handler, err := r.FormFile("image")
    if handler == nil {
        handleErrorResponse(w, &response, "The image file is missing", http.StatusBadRequest)
        return
    }
    if handler.Size > maxFileSizeForUploadImage {
        handleErrorResponse(w, &response, fmt.Sprintf("File size exceeds maximum limit of %d MB", maxFileSizeInMegabytes), http.StatusBadRequest)
        return
    }
	if err != nil {
		handleErrorResponse(w, &response, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()
    originalBuffer := bytes.NewBuffer(nil)
    tee := io.TeeReader(file, originalBuffer)
    compressedBuffer := bytes.NewBuffer(nil)
    cmd := exec.Command("magick", 
        "-",
        "-quality", fmt.Sprint(SignificantCompression),
        "-",
    )
    cmd.Stdin = tee
    cmd.Stdout = compressedBuffer
    if err := cmd.Run(); err != nil {
        handleErrorResponse(w, &response, "Error compressing image", http.StatusInternalServerError)
        return
    }
    response.Base64Image = "data:image/jpeg;base64,"+base64.StdEncoding.EncodeToString(compressedBuffer.Bytes())
    response.OriginalSize = int64(originalBuffer.Len())
    response.CompressedSize = int64(compressedBuffer.Len())
    response.FileName = handler.Filename
    response.Message = "Image compressed successfully"
}
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    response := HealthCheckResponse{}
    response.Alive = true
    defer func() {
        json.NewEncoder(w).Encode(response)
    }()
}

func loadEnvFile() error {
    err := godotenv.Load(".env")
    if err != nil {
        return fmt.Errorf("error loading .env file")
    }
    bearerTokenFromEnvFile = os.Getenv("API_BEARER_TOKEN")
    if bearerTokenFromEnvFile == "" {
        return fmt.Errorf("API_BEARER_TOKEN env is empty")
    }
    return nil
}
func main() {
    errLoadFile := loadEnvFile()
    if errLoadFile != nil {
        log.Fatalf("Failed to load environment: %v", errLoadFile)
    }
	http.HandleFunc("/v1/compress", compressImageHandler)
	http.HandleFunc("/health-check", healthCheckHandler)
	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
