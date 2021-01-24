package output

import (
	"encoding/json"
	"message_api/domain/entity"
	"net/http"
)

// Group .
type Group struct {
	Group *entity.Group `json:"group,omitempty"`
	Error error         `json:"error"`
}

// Groups .
type Groups struct {
	Groups []entity.Group `json:"groups"`
	Error  error          `json:"error"`
}

// WriteResponseJSON .
func (g Group) WriteResponseJSON(w http.ResponseWriter) {
	if g.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		Error{ErrorMessage: g.Error.Error()}.WriteResponseJSON(w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(g)
	}
}

// WriteResponseJSON .
func (g Groups) WriteResponseJSON(w http.ResponseWriter) {
	if g.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		Error{ErrorMessage: g.Error.Error()}.WriteResponseJSON(w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(g)
	}
}
