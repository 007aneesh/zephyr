package main

import (
	"log"
	"net/http"

	"github.com/007aneesh/zephyr/internal/executor"
	"github.com/007aneesh/zephyr/internal/trigger"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Health Check Endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Serverless Platform is running"))
	}).Methods("GET")

	router.HandleFunc("/trigger/http", trigger.HTTPTrigger).Methods("POST")

	log.Println("Deploying a test function...")
	err := executor.DeployFunction("test-function", "./path/to/Dockerfile")
	if err != nil {
		log.Fatalf("Failed to deploy function: %v", err)
	}

	log.Println("Starting HTTP Gateway on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
