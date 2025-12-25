package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func main() {
    // API маршруты
    http.HandleFunc("/api/v1/status", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "status": "success",
            "message": "Go API is running",
        })
    })

    log.Println("Server starting on :8080...")
    // 8080 портында серверді іске қосу
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
