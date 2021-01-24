package application

import (
	"fmt"
	"message_api/domain/input"
	"message_api/domain/output"
	"message_api/repository"
	"time"
)

// CreateGroupMessage .
func CreateGroupMessage(input input.CreateGroupMessageInput) (result output.GroupMessage) {
	groupMessageRepo := repository.GroupMessageRepository{DB: db.DB}
	groupUserRepo := repository.GroupUserRepository{DB: db.DB}

	if !groupUserRepo.IsMember(input.GroupID, input.CurrentUserID) {
		result.Error = fmt.Errorf("no authentication")
		return result
	}

	result.GroupMessage, result.Error = groupMessageRepo.Create(input.GroupID, input.CurrentUserID, input.Text)
	return result
}

// GroupMessageList .
func GroupMessageList(input input.GroupMessageListInput) (result output.GroupMessages) {
	groupMessageRepo := repository.GroupMessageRepository{DB: db.DB}
	groupUserRepo := repository.GroupUserRepository{DB: db.DB}

	if !groupUserRepo.IsMember(input.GroupID, input.CurrentUserID) {
		result.Error = fmt.Errorf("no authentication")
		return result
	}

	result.GroupMessages, result.Error = groupMessageRepo.List(input.GroupID, input.LastMessageID)
	return result
}

// DeleteGroupMessage .
func DeleteGroupMessage(input input.DeleteGroupMessageInput) (result output.Error) {
	groupMessageRepo := repository.GroupMessageRepository{DB: db.DB}

	groupUserRepo := repository.GroupUserRepository{DB: db.DB}

	if !groupUserRepo.IsMember(input.GroupID, input.CurrentUserID) {
		result.ErrorMessage = "no authentication"
		return result
	}

	message, err := groupMessageRepo.Find(input.GroupMessageID)
	if err != nil {
		result.ErrorMessage = err.Error()
		return result
	}

	if message.UserID != input.CurrentUserID {
		result.ErrorMessage = "no authentication"
		return result
	}

	now := time.Now()
	message.DeletedAt = &now
	err = groupMessageRepo.Save(message)

	if err != nil {
		result.ErrorMessage = err.Error()
	}

	return result
}
