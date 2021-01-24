package input

// CreateDirectMessageInput .
type CreateDirectMessageInput struct {
	CurrentUserID   uint
	RecipientUserID uint
	Text            string `json:"text"`
}

// DirectMessageListInput .
type DirectMessageListInput struct {
	CurrentUserID   uint
	RecipientUserID uint
	LastMessageID   uint
}

// DeleteDirectMessageInput .
type DeleteDirectMessageInput struct {
	CurrentUserID   uint
	RecipientUserID uint
	DirectMessageID uint
}
