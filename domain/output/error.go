package output

import (
	"encoding/json"
	"net/http"
)

// Error .
type Error struct {
	ErrorMessage string `json:"error_message"`
}

// WriteResponseJSON .
func (e Error) WriteResponseJSON(w http.ResponseWriter) {
	if e.ErrorMessage == "" {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}
}
