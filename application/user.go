package application

import (
	"message_api/domain/input"
	"message_api/domain/output"
	"message_api/lib/jwt"
	"message_api/repository"
)

// Signup .
func Signup(param input.SignupInput) (result output.User) {
	userRepo := repository.UserRepository{DB: db.DB}

	result.User, result.Error = userRepo.FirstOrCreate(param.Name)
	if result.User != nil {
		result.Token = jwt.CreateToken(result.User.ID)
	}
	return result
}
