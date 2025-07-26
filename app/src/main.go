package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func compressHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("Compress me\n")
	status := map[string]interface{}{
        "service": "image-compressor",
        "status":  "ok",
        "version": "1.0.0",
    }
    json.NewEncoder(w).Encode(status)
}

func main() {
	fmt.Printf("Hello World\n")
	http.HandleFunc("/compress", compressHandler)
	http.ListenAndServe(":8080", nil)
}
