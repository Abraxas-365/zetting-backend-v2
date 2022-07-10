package application

import (
	"service-users/pkg/user/domain/models"
)

func (app *application) Create(new models.User) (bool, error) {

	if can, err := app.userService.CanUserBeCreated(new); !can {
		return false, err
	}
	if _, check, err := app.userRepo.Create(new); !check {
		return false, err
	}
	return true, nil
}
