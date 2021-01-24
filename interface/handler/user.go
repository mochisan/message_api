package handler

import (
	"message_api/application"
	"message_api/domain/input"
	"message_api/domain/output"
	"net/http"
)

// Signup .
func Signup(w http.ResponseWriter, r *http.Request) {
	var param input.SignupInput
	if err := bind(r, &param); err != nil {
		output.Error{ErrorMessage: err.Error()}.WriteResponseJSON(w)
		return
	}
	if param.Name == "" {
		output.Error{ErrorMessage: "name is empty"}.WriteResponseJSON(w)
		return
	}
	application.Signup(param).WriteResponseJSON(w)
}
