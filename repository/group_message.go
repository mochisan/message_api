package repository

import (
	"message_api/domain/entity"

	"github.com/jinzhu/gorm"
)

// GroupMessageRepository .
type GroupMessageRepository struct {
	DB *gorm.DB
}

// List .
func (r GroupMessageRepository) List(groupID, lastMessageID uint) ([]entity.GroupMessage, error) {
	var messages []entity.GroupMessage
	query := r.DB.Where("group_id = ?", groupID).Limit(10).Order("id desc")
	if lastMessageID > 0 {
		query = query.Where("id < ?", lastMessageID)
	}
	err := query.Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// Find .
func (r GroupMessageRepository) Find(id uint) (*entity.GroupMessage, error) {
	var message entity.GroupMessage
	err := r.DB.Find(&message, id).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

// Create .
func (r GroupMessageRepository) Create(groupID, userID uint, text string) (*entity.GroupMessage, error) {
	var message entity.GroupMessage
	message.GroupID = groupID
	message.UserID = userID
	message.Text = text
	err := r.DB.Create(&message).Error
	if err != nil {
		return nil, err
	}
	return r.Find(message.ID)
}

// Save .
func (r GroupMessageRepository) Save(message *entity.GroupMessage) error {
	err := r.DB.Save(message).Error
	return err
}
