package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
        "github.com/Popie52/devboard/backend/internal/handler"
)

func main() {
	fmt.Println("Welcome to DevBoard!")

	mux := handler.RegisterRoutes()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server running on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Println("Failed to start backend:", err)
		os.Exit(1)
	}
}

