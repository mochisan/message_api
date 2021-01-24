package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func init() {

}

func fetchCurrentUserID(r *http.Request) uint {
	return r.Context().Value("current_user_id").(uint)
}

func fetchGroupID(r *http.Request) uint {
	return r.Context().Value("group_id").(uint)
}

func fetchUserID(r *http.Request) uint {
	return r.Context().Value("user_id").(uint)
}

func fetchGroupMessageID(r *http.Request) uint {
	return r.Context().Value("group_message_id").(uint)
}

func fetchDirectMessageID(r *http.Request) uint {
	return r.Context().Value("direct_message_id").(uint)
}

func fetchLastMessageID(r *http.Request) (uint, error) {
	lastMessageIDStr := r.URL.Query().Get("last_message_id")
	lastMessageID, err := strconv.Atoi(lastMessageIDStr)
	if lastMessageIDStr == "" {
		lastMessageID = 0
		err = nil
	}
	return uint(lastMessageID), err
}

func bind(r *http.Request, param interface{}) error {
	return json.NewDecoder(r.Body).Decode(param)
}
