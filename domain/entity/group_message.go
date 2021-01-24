package entity

// GroupMessage .
type GroupMessage struct {
	Model
	UserID  uint   `json:"user_id" gorm:"not null"`
	GroupID uint   `json:"group_id" gorm:"index:idx_message_1; not null"`
	Text    string `json:"text" gorm:"type:text; not null"`

	User User `json:"user"`
}
