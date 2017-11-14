package handlers

import (
	"net/http"
)

// Health responds with 200 while the application is running
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
