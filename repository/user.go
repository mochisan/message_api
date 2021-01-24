package repository

import (
	"message_api/domain/entity"

	"github.com/jinzhu/gorm"
)

// UserRepository .
type UserRepository struct {
	DB *gorm.DB
}

// List .
func (r UserRepository) List() ([]entity.User, error) {
	var users []entity.User
	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Find .
func (r UserRepository) Find(id uint) (*entity.User, error) {
	var user entity.User
	err := r.DB.Find(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Exist .
func (r UserRepository) Exist(id uint) bool {
	var user entity.User
	return !r.DB.Find(&user, id).RecordNotFound()
}

// FirstOrCreate .
func (r UserRepository) FirstOrCreate(name string) (*entity.User, error) {
	var user entity.User
	err := r.DB.FirstOrCreate(&user, entity.User{Name: name}).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Save .
func (r UserRepository) Save(user entity.User) (*entity.User, error) {
	err := r.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
