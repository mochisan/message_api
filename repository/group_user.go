package repository

import (
	"message_api/domain/entity"

	"github.com/jinzhu/gorm"
)

// GroupUserRepository .
type GroupUserRepository struct {
	DB *gorm.DB
}

// Create .
func (r GroupUserRepository) Create(groupID, userID uint) (*entity.GroupUser, error) {
	var groupUser entity.GroupUser
	groupUser.UserID = userID
	groupUser.GroupID = groupID
	err := r.DB.Create(&groupUser).Error
	if err != nil {
		return nil, err
	}
	return &groupUser, nil
}

// IsMember .
func (r GroupUserRepository) IsMember(groupID, userID uint) bool {
	var groupUser entity.GroupUser
	return !r.DB.Table("group_users").Where("user_id = ? and group_id = ?", userID, groupID).First(&groupUser).RecordNotFound()
}
