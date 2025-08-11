package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func compressHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("compressHandler")
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed. Method: " + r.Method, http.StatusMethodNotAllowed)
        return
    }
	// // Run git status command
    // cmd := exec.Command("pwd")
	// output, err := cmd.CombinedOutput()
    // if err != nil {
    //     http.Error(w, fmt.Sprintf("Error running git status: %v", err), http.StatusInternalServerError)
    //     return
    // }
	// cmd2 := exec.Command("cd", "app")
	// output2, err2 := cmd2.CombinedOutput()
    // if err != nil {
    //     http.Error(w, fmt.Sprintf("Error running git status: %v", err2), http.StatusInternalServerError)
    //     return
    // }
	// fmt.Printf("output\n: " + string(output) + "\n")
	// fmt.Printf("output\n: " + string(output2) + "\n")
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("Compress me\n")
	status := map[string]interface{}{
        "service": "image-compressor",
        "status":  "ok",
        "version": "1.0.0",
        "hi": "test measssaa",
    }
    json.NewEncoder(w).Encode(status)
}

func main() {
    http.HandleFunc("/compress", compressHandler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
