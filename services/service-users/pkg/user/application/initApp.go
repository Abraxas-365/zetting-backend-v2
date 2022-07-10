package application

import (
	"service-users/pkg/user/domain/models"
	"service-users/pkg/user/domain/ports"
	"service-users/pkg/user/domain/service"

	"github.com/google/uuid"
)

type Application interface {
	Create(new models.User) (bool, error)
	GetPublicUser(uuid uuid.UUID) (models.UserPublic, bool, error)
	Login(email string, password string) (models.UserPublic, string, error)
}

type application struct {
	userRepo    ports.UserRepo
	userService service.UserService
}

func NewApp(repo ports.UserRepo, service service.UserService) Application {
	return &application{
		repo,
		service,
	}
}
