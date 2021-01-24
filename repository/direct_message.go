package repository

import (
	"message_api/domain/entity"

	"github.com/jinzhu/gorm"
)

// DirectMessageRepository .
type DirectMessageRepository struct {
	DB *gorm.DB
}

// List .
func (r DirectMessageRepository) List(recipientUserID, userID, LastMessageID uint) ([]entity.DirectMessage, error) {
	var messages []entity.DirectMessage
	var err error
	if LastMessageID > 0 {
		err = r.DB.
			Where("id < ? and sender_user_id = ? and recipient_user_id = ?", LastMessageID, recipientUserID, userID).
			Or("id < ? and sender_user_id = ? and recipient_user_id = ?", LastMessageID, userID, recipientUserID).
			Limit(10).Order("id desc").Find(&messages).Error
	} else {
		err = r.DB.
			Where("sender_user_id = ? and recipient_user_id = ?", recipientUserID, userID).
			Or("sender_user_id = ? and recipient_user_id = ?", userID, recipientUserID).
			Limit(10).Order("id desc").Find(&messages).Error
	}
	return messages, err
}

// Find .
func (r DirectMessageRepository) Find(id uint) (*entity.DirectMessage, error) {
	var message entity.DirectMessage
	err := r.DB.Find(&message, id).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

// Create .
func (r DirectMessageRepository) Create(recipientUserID, userID uint, text string) (*entity.DirectMessage, error) {
	var message entity.DirectMessage
	message.RecipientUserID = recipientUserID
	message.SenderUserID = userID
	message.Text = text
	err := r.DB.Create(&message).Error
	if err != nil {
		return nil, err
	}
	return r.Find(message.ID)
}

// Save .
func (r DirectMessageRepository) Save(message *entity.DirectMessage) error {
	err := r.DB.Save(&message).Error
	return err
}
