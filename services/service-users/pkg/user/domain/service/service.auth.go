package service

import (
	"service-users/pkg/user/domain/models"
)

func (r *userService) Auth(email string, password models.Password) (models.User, error) {

	user, check, err := r.userRepo.GetPrivateUser("contact.email", email)
	if !check {
		return models.User{}, err
	}
	if !password.EqualTo(user.Password) {
		return models.User{}, ErrUnauthorized
	}

	return user, nil

}
