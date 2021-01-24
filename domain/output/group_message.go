package output

import (
	"encoding/json"
	"message_api/domain/entity"
	"net/http"
)

// GroupMessage .
type GroupMessage struct {
	GroupMessage *entity.GroupMessage `json:"group_message,omitempty"`
	Error        error                `json:"error"`
}

// GroupMessages .
type GroupMessages struct {
	GroupMessages []entity.GroupMessage `json:"group_messages"`
	Error         error                 `json:"error"`
}

// WriteResponseJSON .
func (g GroupMessage) WriteResponseJSON(w http.ResponseWriter) {
	if g.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		Error{ErrorMessage: g.Error.Error()}.WriteResponseJSON(w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(g)
	}
}

// WriteResponseJSON .
func (g GroupMessages) WriteResponseJSON(w http.ResponseWriter) {
	if g.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		Error{ErrorMessage: g.Error.Error()}.WriteResponseJSON(w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(g)
	}
}
