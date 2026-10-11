package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// buildVersion подставляется при сборке: go build -ldflags "-X main.buildVersion=..."
var buildVersion = "dev"

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{"status": "ok", "version": buildVersion}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	return mux
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("version %s, listening on :%s", buildVersion, port)
	log.Fatal(http.ListenAndServe(":"+port, newMux()))
}
