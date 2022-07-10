package application

import (
	"service-users/internal/auth"
	"service-users/pkg/user/domain/models"
)

func (r *application) Login(email string, password string) (models.UserPublic, string, error) {

	user, err := r.userService.Auth(email, models.Password(password))
	if err != nil {
		return models.UserPublic{}, "", err
	}
	token, err := auth.GereteToken(string(user.Contact.Email), user.ID)
	if err != nil {
		return models.UserPublic{}, "", err
	}

	return user.ToPublic(), token, nil
}
