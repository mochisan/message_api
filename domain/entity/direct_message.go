package entity

// DirectMessage .
type DirectMessage struct {
	Model
	SenderUserID    uint   `json:"sender_user_id" gorm:"index:idx_direct_message_1; not null"`
	RecipientUserID uint   `json:"recipient_user_id" gorm:"index:idx_direct_message_1; not null"`
	Text            string `json:"text" gorm:"type:text; not null"`

	SenderUser    User `json:"sender_user"`
	RecipientUser User `json:"recipient_user"`
}
