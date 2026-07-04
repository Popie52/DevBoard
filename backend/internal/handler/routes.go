package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *NoteHandler) {

	// health
	mux.HandleFunc("/health", HealthHandler)

	// notes collection
	mux.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetNotes(w, r)
		case http.MethodPost:
			h.CreateNote(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// notes item
	mux.HandleFunc("/notes/", h.DeleteNote)
}