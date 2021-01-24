package output

import (
	"encoding/json"
	"message_api/domain/entity"
	"net/http"
)

// User .
type User struct {
	User  *entity.User `json:"user,omitempty"`
	Token string       `json:"token,omitempty"`
	Error error        `json:"error"`
}

// UserList .
type UserList struct {
	Users []entity.User `json:"users"`
	Error error         `json:"error"`
}

// WriteResponseJSON .
func (u User) WriteResponseJSON(w http.ResponseWriter) {
	if u.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		Error{ErrorMessage: u.Error.Error()}.WriteResponseJSON(w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
	}
}

// WriteResponseJSON .
func (u UserList) WriteResponseJSON(w http.ResponseWriter) {
	if u.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		Error{ErrorMessage: u.Error.Error()}.WriteResponseJSON(w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
	}
}
