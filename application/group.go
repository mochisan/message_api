package application

import (
	"message_api/domain/input"
	"message_api/domain/output"
	"message_api/repository"
)

// CreateGroup .
func CreateGroup(input input.CreateGroupInput) (result output.Group) {
	tx := db.DB.Begin()

	groupRepo := repository.GroupRepository{DB: tx}
	groupUserRepo := repository.GroupUserRepository{DB: tx}
	userRepo := repository.UserRepository{DB: tx}

	// グループ作成
	result.Group, result.Error = groupRepo.Create(input.Name)
	if result.Error != nil {
		tx.Rollback()
		return result
	}

	// User全員取得
	users, err := userRepo.List()
	if err != nil {
		result.Error = err
		tx.Rollback()
		return result
	}

	// 全員参加
	for _, user := range users {
		_, result.Error = groupUserRepo.Create(result.Group.ID, user.ID)
		if result.Error != nil {
			tx.Rollback()
			return result
		}
	}

	result.Group.Users = users

	tx.Commit()

	return result
}
