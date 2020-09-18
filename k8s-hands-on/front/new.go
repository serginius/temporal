package main

import (
    "net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    resp := fmt.Sprintf("lahora es %v y hostname %v", time.Now(), os.Getenv("HOSTNAME"))

    w.Write([]byte(resp))
}


func main() {
    http.HandleFunc("/", ServeHTTP)
    http.ListenAndServe(":9090", nil)
}
