package application

import (
	"fmt"
	"message_api/domain/input"
	"message_api/domain/output"
	"message_api/repository"
	"time"
)

// CreateDirectMessage .
func CreateDirectMessage(input input.CreateDirectMessageInput) (result output.DirectMessage) {
	directMessageRepo := repository.DirectMessageRepository{DB: db.DB}
	userRepo := repository.UserRepository{DB: db.DB}

	if !userRepo.Exist(input.RecipientUserID) {
		result.Error = fmt.Errorf("no user")
		return result
	}

	result.DirectMessage, result.Error = directMessageRepo.Create(input.RecipientUserID, input.CurrentUserID, input.Text)
	return result
}

// DirectMessageList .
func DirectMessageList(input input.DirectMessageListInput) (result output.DirectMessages) {
	directMessageRepo := repository.DirectMessageRepository{DB: db.DB}
	userRepo := repository.UserRepository{DB: db.DB}

	if !userRepo.Exist(input.RecipientUserID) {
		result.Error = fmt.Errorf("no user")
		return result
	}

	result.DirectMessages, result.Error = directMessageRepo.List(input.RecipientUserID, input.CurrentUserID, input.LastMessageID)
	return result
}

// DeleteDirectMessage .
func DeleteDirectMessage(input input.DeleteDirectMessageInput) (result output.Error) {
	directMessageRepo := repository.DirectMessageRepository{DB: db.DB}
	userRepo := repository.UserRepository{DB: db.DB}

	if !userRepo.Exist(input.RecipientUserID) {
		result.ErrorMessage = "no user"
		return result
	}

	message, err := directMessageRepo.Find(input.DirectMessageID)
	if err != nil {
		result.ErrorMessage = err.Error()
		return result
	}

	if message.SenderUserID != input.CurrentUserID {
		result.ErrorMessage = "no authentication"
		return result
	}

	if message.RecipientUserID != input.RecipientUserID {
		result.ErrorMessage = "error"
		return result
	}

	now := time.Now()
	message.DeletedAt = &now
	err = directMessageRepo.Save(message)

	if err != nil {
		result.ErrorMessage = err.Error()
	}

	return result
}
