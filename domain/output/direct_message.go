package output

import (
	"encoding/json"
	"message_api/domain/entity"
	"net/http"
)

// DirectMessage .
type DirectMessage struct {
	DirectMessage *entity.DirectMessage `json:"direct_message,omitempty"`
	Error         error                 `json:"error"`
}

// DirectMessages .
type DirectMessages struct {
	DirectMessages []entity.DirectMessage `json:"direct_messages"`
	Error          error                  `json:"error"`
}

// WriteResponseJSON .
func (d DirectMessage) WriteResponseJSON(w http.ResponseWriter) {
	if d.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		Error{ErrorMessage: d.Error.Error()}.WriteResponseJSON(w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(d)
	}
}

// WriteResponseJSON .
func (d DirectMessages) WriteResponseJSON(w http.ResponseWriter) {
	if d.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		Error{ErrorMessage: d.Error.Error()}.WriteResponseJSON(w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(d)
	}
}
