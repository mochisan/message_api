package repository

import (
	"message_api/domain/entity"

	"github.com/jinzhu/gorm"
)

// GroupRepository .
type GroupRepository struct {
	DB *gorm.DB
}

// List .
func (r GroupRepository) List() ([]entity.Group, error) {
	var groups []entity.Group
	err := r.DB.Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// Find .
func (r GroupRepository) Find(id uint) (*entity.Group, error) {
	var group entity.Group
	err := r.DB.Find(&group, id).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// Create .
func (r GroupRepository) Create(name string) (*entity.Group, error) {
	var group entity.Group
	group.Name = name
	err := r.DB.Create(&group).Error
	if err != nil {
		return nil, err
	}
	return r.Find(group.ID)
}

// Save .
func (r GroupRepository) Save(group entity.Group) (*entity.Group, error) {
	err := r.DB.Save(&group).Error
	if err != nil {
		return nil, err
	}
	return r.Find(group.ID)
}
