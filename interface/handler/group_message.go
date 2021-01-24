package handler

import (
	"message_api/application"
	"message_api/domain/input"
	"message_api/domain/output"
	"net/http"
)

// CreateGroupMessage .
func CreateGroupMessage(w http.ResponseWriter, r *http.Request) {
	var input input.CreateGroupMessageInput
	if err := bind(r, &input); err != nil {
		output.Error{ErrorMessage: err.Error()}.WriteResponseJSON(w)
		return
	}

	input.CurrentUserID = fetchCurrentUserID(r)
	input.GroupID = fetchGroupID(r)
	application.CreateGroupMessage(input).WriteResponseJSON(w)
}

// GroupMessageList .
func GroupMessageList(w http.ResponseWriter, r *http.Request) {
	lastMessageID, err := fetchLastMessageID(r)
	if err != nil {
		output.Error{ErrorMessage: err.Error()}.WriteResponseJSON(w)
		return
	}

	var input = input.GroupMessageListInput{
		LastMessageID: lastMessageID,
		CurrentUserID: fetchCurrentUserID(r),
		GroupID:       fetchGroupID(r),
	}
	application.GroupMessageList(input).WriteResponseJSON(w)
}

// DeleteGroupMessage .
func DeleteGroupMessage(w http.ResponseWriter, r *http.Request) {
	var input = input.DeleteGroupMessageInput{
		CurrentUserID:  fetchCurrentUserID(r),
		GroupMessageID: fetchGroupMessageID(r),
		GroupID:        fetchGroupID(r),
	}
	application.DeleteGroupMessage(input).WriteResponseJSON(w)
}
