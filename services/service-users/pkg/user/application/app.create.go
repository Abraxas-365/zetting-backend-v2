package application

import (
	"service-users/pkg/user/domain/models"
)

func (app *application) Create(new models.User) (bool, error) {

	if can, err := app.userService.CanUserBeCreated(new); !can {
		return false, err
	}
	if event, worked, err := app.userRepo.Create(new); !worked {
		return false, err
	} else {
		if err := app.mqpublisher.PublishEvent(event); err != nil {
			return false, err
		}
	}

	return true, nil
}
