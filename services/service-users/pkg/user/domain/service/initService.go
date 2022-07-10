package service

import (
	"errors"
	"service-users/pkg/user/domain/models"
	"service-users/pkg/user/domain/ports"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
	ErrEmptyParams  = errors.New("Empty parameters")
	ErrUserExists   = errors.New("User already exists")
	ErrBadPassword  = errors.New("Bad Password Param")
)

type UserService interface {
	Auth(email string, password models.Password) (models.User, error)
	CanUserBeCreated(user models.User) (bool, error)
}

type userService struct {
	userRepo ports.UserRepo
}

func NewUserService(userRepo ports.UserRepo) UserService {
	return &userService{
		userRepo,
	}
}
