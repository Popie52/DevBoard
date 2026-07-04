package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Popie52/devboard/backend/internal/handler"
	"github.com/Popie52/devboard/backend/internal/middleware"
	"github.com/Popie52/devboard/backend/internal/service"
)

func main() {

	// service
	noteService := service.NewNoteService()

	// handler
	noteHandler := handler.NewNoteHandler(noteService)

	// router
	mux := http.NewServeMux()

	// routes
	handler.RegisterRoutes(mux, noteHandler)

	// middleware chain
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
		log.Println("Server failed:", err)
		os.Exit(1)
	}
}