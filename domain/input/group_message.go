package input

// CreateGroupMessageInput .
type CreateGroupMessageInput struct {
	CurrentUserID uint
	GroupID       uint
	Text          string `json:"text"`
}

// GroupMessageListInput .
type GroupMessageListInput struct {
	CurrentUserID uint
	GroupID       uint
	LastMessageID uint
}

// DeleteGroupMessageInput .
type DeleteGroupMessageInput struct {
	CurrentUserID  uint
	GroupID        uint
	GroupMessageID uint
}
