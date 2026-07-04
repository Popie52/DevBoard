package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
    
	"github.com/Popie52/devboard/backend/internal/handler"
	"github.com/Popie52/devboard/backend/internal/middleware"	
)

func main() {
	fmt.Println("Welcome to DevBoard!")

	mux := handler.RegisterRoutes()

	h := middleware.Logger(mux)
	h = middleware.CORSHANDLER(h)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      h,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server running on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Println("Failed to start backend:", err)
		os.Exit(1)
	}
}

