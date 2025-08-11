package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const maxFileSizeForUploadImage int64 = 10 * 1024 * 1024 // 10 MB
// TODO: I want to move this to a utils directory but where should I move this?
type CompressionQuality int
const (
    LosslessCompression CompressionQuality = 100
	VeryHighCompression CompressionQuality = 95
	WebOptimizedCompression CompressionQuality = 85
	ModerateCompression CompressionQuality = 75
	SignificantCompression CompressionQuality = 60
	ExtremeCompression CompressionQuality = 10
)
func compressHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("perry: compressHandler")
	if r.Method != http.MethodPost {
        // TODO: 成功と失敗の時にJSONを返却するように修正
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(maxFileSizeForUploadImage)
	if err != nil {
        // TODO: 成功と失敗の時にJSONを返却するように修正
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("image")
	if err != nil {
        // TODO: 成功と失敗の時にJSONを返却するように修正
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	tempDir := "./tmp"
	os.MkdirAll(tempDir, os.ModePerm)

	// TODO: レースコンディションになってるっぽい。複数のユーザが同じファイル名のファイルが渡されたらバグりそう
	originalPath := filepath.Join(tempDir, "original_"+handler.Filename)
	compressedPath := filepath.Join(tempDir, "compressed_"+handler.Filename)
	originalFile, err := os.Create(originalPath)
	if err != nil {
        // TODO: 成功と失敗の時にJSONを返却するように修正
		http.Error(w, "Error creating original file", http.StatusInternalServerError)
		return
	}
	defer originalFile.Close()
    _, err = io.Copy(originalFile, file)
	if err != nil {
        // TODO: 成功と失敗の時にJSONを返却するように修正
		http.Error(w, "Error saving original file", http.StatusInternalServerError)
		return
	}
	cmd := exec.Command("magick", 
		originalPath,
		"-quality", fmt.Sprint(SignificantCompression),  // compression quality (0-100, lower is more compressed)
		compressedPath,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
        // TODO: 成功と失敗の時にJSONを返却するように修正
		log.Printf("Compression error: %v, Output: %s", err, string(output))
		http.Error(w, "Error compressing image", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/"+filepath.Ext(handler.Filename)[1:])
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(compressedPath))
	compressedFile, err := os.Open(compressedPath)
	if err != nil {
        // TODO: 成功と失敗の時にJSONを返却するように修正
		http.Error(w, "Error reading compressed file", http.StatusInternalServerError)
		return
	}
	defer compressedFile.Close()
	defer os.Remove(originalPath)
	defer os.Remove(compressedPath)
    // 圧縮された画像をレスポンスにコピー
	_, err = io.Copy(w, compressedFile)
	if err != nil {
		log.Printf("Error serving compressed file: %v", err)
	}
}

func main() {
	http.HandleFunc("/compress", compressHandler)
	// tmpディレクトリ存在を確認
	os.MkdirAll("./tmp", os.ModePerm)
	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
