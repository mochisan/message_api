package handler

import (
	"encoding/json"
	"message_api/application"
	"message_api/domain/input"
	"message_api/domain/output"
	"net/http"
)

// CreateGroup .
func CreateGroup(w http.ResponseWriter, r *http.Request) {
	var param input.CreateGroupInput
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		output.Error{ErrorMessage: err.Error()}.WriteResponseJSON(w)
		return
	}
	param.CurrentUserID = fetchCurrentUserID(r)
	application.CreateGroup(param).WriteResponseJSON(w)
}
