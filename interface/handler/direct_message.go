package handler

import (
	"message_api/application"
	"message_api/domain/input"
	"message_api/domain/output"
	"net/http"
)

// CreateDirectMessage .
func CreateDirectMessage(w http.ResponseWriter, r *http.Request) {
	var input input.CreateDirectMessageInput
	if err := bind(r, &input); err != nil {
		output.Error{ErrorMessage: err.Error()}.WriteResponseJSON(w)
		return
	}

	input.CurrentUserID = fetchCurrentUserID(r)
	input.RecipientUserID = fetchUserID(r)
	application.CreateDirectMessage(input).WriteResponseJSON(w)
}

// DirectMessageList .
func DirectMessageList(w http.ResponseWriter, r *http.Request) {
	lastMessageID, err := fetchLastMessageID(r)
	if err != nil {
		output.Error{ErrorMessage: err.Error()}.WriteResponseJSON(w)
		return
	}

	var input = input.DirectMessageListInput{
		LastMessageID:   lastMessageID,
		CurrentUserID:   fetchCurrentUserID(r),
		RecipientUserID: fetchUserID(r),
	}
	application.DirectMessageList(input).WriteResponseJSON(w)
}

// DeleteDirectMessage .
func DeleteDirectMessage(w http.ResponseWriter, r *http.Request) {
	var input = input.DeleteDirectMessageInput{
		CurrentUserID:   fetchCurrentUserID(r),
		DirectMessageID: fetchDirectMessageID(r),
		RecipientUserID: fetchUserID(r),
	}
	application.DeleteDirectMessage(input).WriteResponseJSON(w)
}
